// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppCoinsColumns holds the columns for the "app_coins" table.
	AppCoinsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "display_names", Type: field.TypeJSON, Nullable: true},
		{Name: "logo", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "for_pay", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "withdraw_auto_review_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "product_page", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "daily_reward_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "display", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "display_index", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "max_amount_per_withdraw", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// AppCoinsTable holds the schema information for the "app_coins" table.
	AppCoinsTable = &schema.Table{
		Name:       "app_coins",
		Columns:    AppCoinsColumns,
		PrimaryKey: []*schema.Column{AppCoinsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appcoin_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppCoinsColumns[4]},
			},
		},
	}
	// ChainBasesColumns holds the columns for the "chain_bases" table.
	ChainBasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "logo", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "native_unit", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "atomic_unit", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "unit_exp", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "env", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "chain_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "nickname", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "gas_type", Type: field.TypeString, Nullable: true, Default: "DefaultGasType"},
	}
	// ChainBasesTable holds the schema information for the "chain_bases" table.
	ChainBasesTable = &schema.Table{
		Name:       "chain_bases",
		Columns:    ChainBasesColumns,
		PrimaryKey: []*schema.Column{ChainBasesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "chainbase_ent_id",
				Unique:  true,
				Columns: []*schema.Column{ChainBasesColumns[4]},
			},
		},
	}
	// CoinBasesColumns holds the columns for the "coin_bases" table.
	CoinBasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "logo", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "presale", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "unit", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "env", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "reserved_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "for_pay", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// CoinBasesTable holds the schema information for the "coin_bases" table.
	CoinBasesTable = &schema.Table{
		Name:       "coin_bases",
		Columns:    CoinBasesColumns,
		PrimaryKey: []*schema.Column{CoinBasesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "coinbase_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CoinBasesColumns[4]},
			},
		},
	}
	// CoinDescriptionsColumns holds the columns for the "coin_descriptions" table.
	CoinDescriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "title", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// CoinDescriptionsTable holds the schema information for the "coin_descriptions" table.
	CoinDescriptionsTable = &schema.Table{
		Name:       "coin_descriptions",
		Columns:    CoinDescriptionsColumns,
		PrimaryKey: []*schema.Column{CoinDescriptionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "coindescription_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CoinDescriptionsColumns[4]},
			},
		},
	}
	// CoinExtrasColumns holds the columns for the "coin_extras" table.
	CoinExtrasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "home_page", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "specs", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "stable_usd", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// CoinExtrasTable holds the schema information for the "coin_extras" table.
	CoinExtrasTable = &schema.Table{
		Name:       "coin_extras",
		Columns:    CoinExtrasColumns,
		PrimaryKey: []*schema.Column{CoinExtrasColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "coinextra_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CoinExtrasColumns[4]},
			},
		},
	}
	// CoinFiatsColumns holds the columns for the "coin_fiats" table.
	CoinFiatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "fiat_id", Type: field.TypeUUID, Nullable: true},
		{Name: "feed_type", Type: field.TypeString, Nullable: true, Default: "DefaultFeedType"},
	}
	// CoinFiatsTable holds the schema information for the "coin_fiats" table.
	CoinFiatsTable = &schema.Table{
		Name:       "coin_fiats",
		Columns:    CoinFiatsColumns,
		PrimaryKey: []*schema.Column{CoinFiatsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "coinfiat_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CoinFiatsColumns[4]},
			},
		},
	}
	// CoinFiatCurrenciesColumns holds the columns for the "coin_fiat_currencies" table.
	CoinFiatCurrenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "fiat_id", Type: field.TypeUUID, Nullable: true},
		{Name: "feed_type", Type: field.TypeString, Nullable: true, Default: "DefaultFeedType"},
		{Name: "market_value_low", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "market_value_high", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// CoinFiatCurrenciesTable holds the schema information for the "coin_fiat_currencies" table.
	CoinFiatCurrenciesTable = &schema.Table{
		Name:       "coin_fiat_currencies",
		Columns:    CoinFiatCurrenciesColumns,
		PrimaryKey: []*schema.Column{CoinFiatCurrenciesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "coinfiatcurrency_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CoinFiatCurrenciesColumns[4]},
			},
		},
	}
	// CoinFiatCurrencyHistoriesColumns holds the columns for the "coin_fiat_currency_histories" table.
	CoinFiatCurrencyHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "fiat_id", Type: field.TypeUUID, Nullable: true},
		{Name: "feed_type", Type: field.TypeString, Nullable: true, Default: "DefaultFeedType"},
		{Name: "market_value_low", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "market_value_high", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// CoinFiatCurrencyHistoriesTable holds the schema information for the "coin_fiat_currency_histories" table.
	CoinFiatCurrencyHistoriesTable = &schema.Table{
		Name:       "coin_fiat_currency_histories",
		Columns:    CoinFiatCurrencyHistoriesColumns,
		PrimaryKey: []*schema.Column{CoinFiatCurrencyHistoriesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "coinfiatcurrencyhistory_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CoinFiatCurrencyHistoriesColumns[4]},
			},
		},
	}
	// CurrenciesColumns holds the columns for the "currencies" table.
	CurrenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "feed_type", Type: field.TypeString, Nullable: true, Default: "DefaultFeedType"},
		{Name: "market_value_high", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "market_value_low", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// CurrenciesTable holds the schema information for the "currencies" table.
	CurrenciesTable = &schema.Table{
		Name:       "currencies",
		Columns:    CurrenciesColumns,
		PrimaryKey: []*schema.Column{CurrenciesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "currency_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CurrenciesColumns[4]},
			},
			{
				Name:    "currency_coin_type_id_id",
				Unique:  false,
				Columns: []*schema.Column{CurrenciesColumns[5], CurrenciesColumns[0]},
			},
			{
				Name:    "currency_coin_type_id",
				Unique:  false,
				Columns: []*schema.Column{CurrenciesColumns[5]},
			},
		},
	}
	// CurrencyFeedsColumns holds the columns for the "currency_feeds" table.
	CurrencyFeedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "feed_type", Type: field.TypeString, Nullable: true, Default: "DefaultFeedType"},
		{Name: "feed_coin_name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// CurrencyFeedsTable holds the schema information for the "currency_feeds" table.
	CurrencyFeedsTable = &schema.Table{
		Name:       "currency_feeds",
		Columns:    CurrencyFeedsColumns,
		PrimaryKey: []*schema.Column{CurrencyFeedsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "currencyfeed_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CurrencyFeedsColumns[4]},
			},
			{
				Name:    "currencyfeed_coin_type_id_id",
				Unique:  false,
				Columns: []*schema.Column{CurrencyFeedsColumns[5], CurrencyFeedsColumns[0]},
			},
			{
				Name:    "currencyfeed_coin_type_id",
				Unique:  false,
				Columns: []*schema.Column{CurrencyFeedsColumns[5]},
			},
		},
	}
	// CurrencyHistoriesColumns holds the columns for the "currency_histories" table.
	CurrencyHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "feed_type", Type: field.TypeString, Nullable: true, Default: "DefaultFeedType"},
		{Name: "market_value_high", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "market_value_low", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// CurrencyHistoriesTable holds the schema information for the "currency_histories" table.
	CurrencyHistoriesTable = &schema.Table{
		Name:       "currency_histories",
		Columns:    CurrencyHistoriesColumns,
		PrimaryKey: []*schema.Column{CurrencyHistoriesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "currencyhistory_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CurrencyHistoriesColumns[4]},
			},
			{
				Name:    "currencyhistory_coin_type_id_id",
				Unique:  false,
				Columns: []*schema.Column{CurrencyHistoriesColumns[5], CurrencyHistoriesColumns[0]},
			},
			{
				Name:    "currencyhistory_coin_type_id",
				Unique:  false,
				Columns: []*schema.Column{CurrencyHistoriesColumns[5]},
			},
		},
	}
	// ExchangeRatesColumns holds the columns for the "exchange_rates" table.
	ExchangeRatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "market_value", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "settle_value", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "settle_percent", Type: field.TypeUint32, Nullable: true, Default: 100},
		{Name: "settle_tips", Type: field.TypeJSON, Nullable: true},
		{Name: "setter", Type: field.TypeUUID, Nullable: true},
	}
	// ExchangeRatesTable holds the schema information for the "exchange_rates" table.
	ExchangeRatesTable = &schema.Table{
		Name:       "exchange_rates",
		Columns:    ExchangeRatesColumns,
		PrimaryKey: []*schema.Column{ExchangeRatesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "exchangerate_ent_id",
				Unique:  true,
				Columns: []*schema.Column{ExchangeRatesColumns[4]},
			},
		},
	}
	// FiatsColumns holds the columns for the "fiats" table.
	FiatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "logo", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "unit", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// FiatsTable holds the schema information for the "fiats" table.
	FiatsTable = &schema.Table{
		Name:       "fiats",
		Columns:    FiatsColumns,
		PrimaryKey: []*schema.Column{FiatsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "fiat_ent_id",
				Unique:  true,
				Columns: []*schema.Column{FiatsColumns[4]},
			},
		},
	}
	// FiatCurrenciesColumns holds the columns for the "fiat_currencies" table.
	FiatCurrenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "fiat_id", Type: field.TypeUUID, Nullable: true},
		{Name: "feed_type", Type: field.TypeString, Nullable: true, Default: "DefaultFeedType"},
		{Name: "market_value_low", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "market_value_high", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// FiatCurrenciesTable holds the schema information for the "fiat_currencies" table.
	FiatCurrenciesTable = &schema.Table{
		Name:       "fiat_currencies",
		Columns:    FiatCurrenciesColumns,
		PrimaryKey: []*schema.Column{FiatCurrenciesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "fiatcurrency_ent_id",
				Unique:  true,
				Columns: []*schema.Column{FiatCurrenciesColumns[4]},
			},
		},
	}
	// FiatCurrencyFeedsColumns holds the columns for the "fiat_currency_feeds" table.
	FiatCurrencyFeedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "fiat_id", Type: field.TypeUUID, Nullable: true},
		{Name: "feed_type", Type: field.TypeString, Nullable: true, Default: "DefaultFeedType"},
		{Name: "feed_fiat_name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// FiatCurrencyFeedsTable holds the schema information for the "fiat_currency_feeds" table.
	FiatCurrencyFeedsTable = &schema.Table{
		Name:       "fiat_currency_feeds",
		Columns:    FiatCurrencyFeedsColumns,
		PrimaryKey: []*schema.Column{FiatCurrencyFeedsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "fiatcurrencyfeed_ent_id",
				Unique:  true,
				Columns: []*schema.Column{FiatCurrencyFeedsColumns[4]},
			},
			{
				Name:    "fiatcurrencyfeed_fiat_id_id",
				Unique:  false,
				Columns: []*schema.Column{FiatCurrencyFeedsColumns[5], FiatCurrencyFeedsColumns[0]},
			},
			{
				Name:    "fiatcurrencyfeed_fiat_id",
				Unique:  false,
				Columns: []*schema.Column{FiatCurrencyFeedsColumns[5]},
			},
		},
	}
	// FiatCurrencyHistoriesColumns holds the columns for the "fiat_currency_histories" table.
	FiatCurrencyHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "fiat_id", Type: field.TypeUUID, Nullable: true},
		{Name: "feed_type", Type: field.TypeString, Nullable: true, Default: "DefaultFeedType"},
		{Name: "market_value_low", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "market_value_high", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// FiatCurrencyHistoriesTable holds the schema information for the "fiat_currency_histories" table.
	FiatCurrencyHistoriesTable = &schema.Table{
		Name:       "fiat_currency_histories",
		Columns:    FiatCurrencyHistoriesColumns,
		PrimaryKey: []*schema.Column{FiatCurrencyHistoriesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "fiatcurrencyhistory_ent_id",
				Unique:  true,
				Columns: []*schema.Column{FiatCurrencyHistoriesColumns[4]},
			},
		},
	}
	// SettingsColumns holds the columns for the "settings" table.
	SettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "fee_coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "withdraw_fee_by_stable_usd", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "withdraw_fee_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "collect_fee_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "hot_wallet_fee_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "low_fee_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "hot_low_fee_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "hot_wallet_account_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "payment_account_collect_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "least_transfer_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "need_memo", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "refresh_currency", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "check_new_address_balance", Type: field.TypeBool, Nullable: true, Default: true},
	}
	// SettingsTable holds the schema information for the "settings" table.
	SettingsTable = &schema.Table{
		Name:       "settings",
		Columns:    SettingsColumns,
		PrimaryKey: []*schema.Column{SettingsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "setting_ent_id",
				Unique:  true,
				Columns: []*schema.Column{SettingsColumns[4]},
			},
		},
	}
	// TransColumns holds the columns for the "trans" table.
	TransColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "from_account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "to_account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "fee_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "chain_tx_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "DefaultTxState"},
		{Name: "extra", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "type", Type: field.TypeString, Nullable: true, Default: "DefaultTxType"},
	}
	// TransTable holds the schema information for the "trans" table.
	TransTable = &schema.Table{
		Name:       "trans",
		Columns:    TransColumns,
		PrimaryKey: []*schema.Column{TransColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "tran_ent_id",
				Unique:  true,
				Columns: []*schema.Column{TransColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppCoinsTable,
		ChainBasesTable,
		CoinBasesTable,
		CoinDescriptionsTable,
		CoinExtrasTable,
		CoinFiatsTable,
		CoinFiatCurrenciesTable,
		CoinFiatCurrencyHistoriesTable,
		CurrenciesTable,
		CurrencyFeedsTable,
		CurrencyHistoriesTable,
		ExchangeRatesTable,
		FiatsTable,
		FiatCurrenciesTable,
		FiatCurrencyFeedsTable,
		FiatCurrencyHistoriesTable,
		SettingsTable,
		TransTable,
	}
)

func init() {
}
