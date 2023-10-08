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
)

const (
	keyUsername  = "username"
	keyPassword  = "password"
	keyDBName    = "database_name"
	maxOpen      = 10
	maxIdle      = 10
	MaxLife      = 3
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

func tables(ctx context.Context, dbName string) ([]string, error) {
	tables := []string{}
	rows, err := tx.QueryContext(
		ctx,
		fmt.Sprintf("select table_name from information_schema.tables where table_schema = '%v'", dbName),
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&tables); err != nil {
			return err
		}
	}
	return nil
}

func migrateEntID(ctx context.Context, dbName, table string, tx *sql.Tx) error {
	logger.Sugar().Infow(
		"migrateEntID",
		"db", dbName,
		"table", table,
	)

	rc := 0
	rows, err := tx.QueryContext(
		ctx,
		fmt.Sprintf("select 1 from information_schema.columns where table_schema='%v' and table_name='%v' and column_name='ent_id'", dbName, table),
	)
	if err != nil {
		return err
	}
	for rows.Next() {
		if err := rows.Scan(&rc); err != nil {
			return err
		}
	}
	if rc != 0 {
		return nil
	}
	_type := ""
	rows, err = tx.QueryContext(
		ctx,
		fmt.Sprintf("select data_type from information_schema.columns where table_name='%v' and column_name='id' and table_schema='%v'", table, dbName),
	)
	if err != nil {
		return err
	}
	for rows.Next() {
		if err := rows.Scan(&_type); err != nil {
			return err
		}
	}
	if strings.Contains(_type, "int") {
		logger.Sugar().Infow(
			"migrateEntID",
			"db", dbName,
			"table", table,
			"State", "Migrated",
		)
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
		fmt.Sprintf("alter table %v.%v change column id ent_id char(36)", dbName, table),
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

	logger.Sugar().Infow("Migrate order", "Start", "...")
	defer func(err *error) {
		_ = redis2.Unlock(lockKey())
		logger.Sugar().Infow("Migrate order", "Done", "...", "error", *err)
	}(&err)

	err = redis2.TryLock(lockKey(), 0)
	if err != nil {
		return err
	}

	conn, err = open(servicename.ServiceDomain)
	if err != nil {
		return err
	}
	defer conn.Close()

	tx, err = conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	_tables, err := tables(ctx, "chain_manager")
	if err != nil {
		return err
	}

	for _, table := range tables {
		if err = migrateEntID(ctx, table, tx); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	_ = tx.Commit()
	return nil
}
