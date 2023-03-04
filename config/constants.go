package config

const (
	AppName     = "bpjs"
	AppDLayout  = "2006-01-02"
	AppTLayout  = "2006-01-02 15:04:05"
	AppTimeZone = "Asia/Jakarta"

	CorporateActionEventIDDividend     = 1
	CorporateActionEventIDStockSplit   = 2
	CorporateActionEventIDReverseSplit = 3
	CorporateActionEventIDRightIssue   = 4
	CorporateActionEventIDWarrant      = 5
	CorporateActionEventIDBonus        = 6
	CorporateActionEventIDRUPS         = 7
	CorporateActionEventIDPublicExpose = 8
	CorporateActionEventIDIPO          = 9

	LimitStockSearchResult = 25
	LimitStockSearchQuery  = 2000

	SearchErrCodeDuplicateEntry = "1062"
	SearchErrCodeDataNotFound   = "1452"

	RedisTTLOneHour         = 1 * 1 * 1 * 3600 * 1000 * 1000 * 1000
	RedisTTLOneDay          = 1 * 24 * 1 * 3600 * 1000 * 1000 * 1000
	RedisTTLOneWeek         = 7 * 24 * 1 * 3600 * 1000 * 1000 * 1000
	RedisErrKeyDoesNotExist = "Key does not exist in Redis"

	RedisKeyIndices                   = "indices"
	RedisKeyCompanyFinancialAnnually  = "stock:company-financial:annually"
	RedisKeyCompanyFinancialQuarterly = "stock:company-financial:quarterly"
	RedisKeyStockDetail               = "stock:detail"
	RedisKeyStockKeyStat              = "stock:keystat"
	RedisKeyStockSummary              = "stock:summary"
	RedisKeyStockTrades               = "stock:%s:rt"
	RedisKeyTrades                    = "rt"

	RedisStartTrades = 0
	RedisLimitTrades = 19

	StockSummaryOneMonth    = "1m"
	StockSummaryThreeMonths = "3m"
	StockSummaryYearToDate  = "ytd"
	StockSummaryOneYear     = "1y"
	StockSummaryThreeYears  = "3y"
	StockSummaryFiveYears   = "5y"
)
