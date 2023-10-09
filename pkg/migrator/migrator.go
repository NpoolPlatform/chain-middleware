//nolint:nolintlint
package migrator

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	servicename "github.com/NpoolPlatform/chain-middleware/pkg/servicename"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	constant "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

const (
	keyUsername  = "username"
	keyPassword  = "password"
	keyDBName    = "database_name"
	maxOpen      = 2
	maxIdle      = 1
	MaxLife      = 0
	keyServiceID = "serviceid"
)

func lockKey() string {
	serviceID := config.GetStringValueWithNameSpace(servicename.ServiceDomain, keyServiceID)
	return fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixMigrate, serviceID)
}

func dsn(hostname string) (string, error) {
	username := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyUsername)
	password := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyPassword)
	dbname := config.GetStringValueWithNameSpace(hostname, keyDBName)

	svc, err := config.PeekService(constant.MysqlServiceName)
	if err != nil {
		logger.Sugar().Warnw("dsn", "error", err)
		return "", err
	}

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&interpolateParams=true",
		username, password,
		svc.Address,
		svc.Port,
		dbname,
	), nil
}

func open(hostname string) (conn *sql.DB, err error) {
	hdsn, err := dsn(hostname)
	if err != nil {
		return nil, err
	}

	logger.Sugar().Infow("open", "hdsn", hdsn)

	conn, err = sql.Open("mysql", hdsn)
	if err != nil {
		return nil, err
	}

	// https://github.com/go-sql-driver/mysql
	// See "Important settings" section.

	conn.SetConnMaxLifetime(time.Minute * MaxLife)
	conn.SetMaxOpenConns(maxOpen)
	conn.SetMaxIdleConns(maxIdle)

	return conn, nil
}

func tables(ctx context.Context, dbName string, tx *sql.Tx) ([]string, error) {
	tables := []string{}
	rows, err := tx.QueryContext(
		ctx,
		fmt.Sprintf("select table_name from information_schema.tables where table_schema = '%v'", dbName),
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		table := []byte{}
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, string(table))
	}
	logger.Sugar().Infow(
		"tables",
		"Tables", tables,
	)
	return tables, nil
}

func migrateEntID(ctx context.Context, dbName, table string, tx *sql.Tx) error {
	logger.Sugar().Infow(
		"migrateEntID",
		"db", dbName,
		"table", table,
	)

	_type := []byte{}
	rows, err := tx.QueryContext(
		ctx,
		fmt.Sprintf("select column_type from information_schema.columns where table_name='%v' and column_name='id' and table_schema='%v'", table, dbName),
	)
	if err != nil {
		return err
	}
	for rows.Next() {
		if err := rows.Scan(&_type); err != nil {
			return err
		}
	}
	if strings.Contains(string(_type), "int") && !strings.Contains(string(_type), "unsigned") {
		logger.Sugar().Infow(
			"migrateEntID",
			"db", dbName,
			"table", table,
			"type", string(_type),
			"State", "INT UNSIGNED",
		)
		_, err = tx.ExecContext(
			ctx,
			fmt.Sprintf("alter table %v.%v change id id int unsigned not null auto_increment", dbName, table),
		)
		if err != nil {
			return err
		}
	}
	key := ""
	rc := 0
	rows, err = tx.QueryContext(
		ctx,
		fmt.Sprintf("select column_key,1 from information_schema.columns where table_schema='%v' and table_name='%v' and column_name='ent_id'", dbName, table),
	)
	if err != nil {
		return err
	}
	for rows.Next() {
		if err := rows.Scan(&key, &rc); err != nil {
			return err
		}
	}
	if rc == 1 {
		if key != "UNI" {
			logger.Sugar().Infow(
				"migrateEntID",
				"db", dbName,
				"table", table,
				"State", "ENT_ID UNIQUE",
			)
			_, err = tx.ExecContext(
				ctx,
				fmt.Sprintf("alter table %v.%v change column ent_id ent_id char(36) unique", dbName, table),
			)
			if err != nil {
				return err
			}
		}
		logger.Sugar().Infow(
			"migrateEntID",
			"db", dbName,
			"table", table,
			"State", "ENT_ID UUID",
		)
		rows, err := tx.QueryContext(
			ctx,
			fmt.Sprintf("select id from %v.%v where ent_id=''", dbName, table),
		)
		if err != nil {
			return err
		}
		for rows.Next() {
			var id uint32
			if err := rows.Scan(&id); err != nil {
				return err
			}
			if _, err := tx.ExecContext(
				ctx,
				fmt.Sprintf("update %v.%v set ent_id='%v' where id=%v", dbName, table, uuid.New(), id),
			); err != nil {
				return err
			}
		}
		return nil
	}
	logger.Sugar().Infow(
		"migrateEntID",
		"db", dbName,
		"table", table,
		"State", "ID -> EntID",
	)
	_, err = tx.ExecContext(
		ctx,
		fmt.Sprintf("alter table %v.%v change column id ent_id char(36) unique", dbName, table),
	)
	if err != nil {
		return err
	}
	logger.Sugar().Infow(
		"migrateEntID",
		"db", dbName,
		"table", table,
		"State", "ID INT",
	)
	_, err = tx.ExecContext(
		ctx,
		fmt.Sprintf("alter table %v.%v add id int unsigned not null auto_increment, drop primary key, add primary key(id)", dbName, table),
	)
	if err != nil {
		return err
	}
	rows, err = tx.QueryContext(
		ctx,
		fmt.Sprintf("select id from %v.%v where ent_id=''", dbName, table),
	)
	if err != nil {
		return err
	}
	logger.Sugar().Infow(
		"migrateEntID",
		"db", dbName,
		"table", table,
		"State", "ENT_ID UUID",
	)
	for rows.Next() {
		var id uint32
		if err := rows.Scan(&id); err != nil {
			return err
		}
		if _, err := tx.ExecContext(
			ctx,
			fmt.Sprintf("update %v.%v set ent_id='%v' where id=%v", dbName, table, uuid.New(), id),
		); err != nil {
			return err
		}
	}
	logger.Sugar().Infow(
		"migrateEntID",
		"db", dbName,
		"table", table,
		"State", "Migrated",
	)
	return err
}

func Migrate(ctx context.Context) error {
	var err error
	var conn *sql.DB
	var tx *sql.Tx

	logger.Sugar().Infow("Migrate", "Start", "...")
	defer func(err *error) {
		_ = redis2.Unlock(lockKey())
		logger.Sugar().Infow("Migrate", "Done", "...", "error", *err)
	}(&err)

	err = redis2.TryLock(lockKey(), 0)
	if err != nil {
		return err
	}

	conn, err = open(servicename.ServiceDomain)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Sugar().Errorw("Close", "Error", err)
		}
	}()

	tx, err = conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	dbname := config.GetStringValueWithNameSpace(servicename.ServiceDomain, keyDBName)
	_tables, err := tables(ctx, dbname, tx)
	if err != nil {
		return err
	}

	for _, table := range _tables {
		if err = migrateEntID(ctx, dbname, table, tx); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	_ = tx.Commit()
	return nil
}
