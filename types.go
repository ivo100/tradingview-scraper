package tradingview

// SocketInterface ...
type SocketInterface interface {
	AddSymbol(symbol string) error
	RemoveSymbol(symbol string) error
	Init(fields ...string) error
	Close() error
}

// SocketMessage ...
type SocketMessage struct {
	Message string      `json:"m"`
	Payload interface{} `json:"p"`
}

// QuoteMessage ...
type QuoteMessage struct {
	Symbol string     `mapstructure:"n"`
	Status string     `mapstructure:"s"`
	Data   *QuoteData `mapstructure:"v"`
}

//	getSocketMessage("quote_set_fields", []string{s.sessionID, "lp", "volume", "bid", "ask", "ch", "chp"}),
//
// QuoteData ...
type QuoteData struct {
	Price             *float64 `mapstructure:"lp"`
	PrevClosePrice    *float64 `mapstructure:"prev_close_price"`
	RegularClosePrice *float64 `mapstructure:"regular_close_price"`
	RegularCloseTime  *int64   `mapstructure:"regular_close_time"`
	HighPrice         *float64 `mapstructure:"high_price"`
	LowPrice          *float64 `mapstructure:"low_price"`
	OpenPrice         *float64 `mapstructure:"open_price"`
	OpenTime          *int64   `mapstructure:"open_time"`
	Volume            *float64 `mapstructure:"volume"`
	Bid               *float64 `mapstructure:"bid"`
	Ask               *float64 `mapstructure:"ask"`
	Change            *float64 `mapstructure:"ch"`
	Time              *int64   `mapstructure:"lp_time"`
}

// Flags ...
type Flags struct {
	Flags []string `json:"flags"`
}

// OnReceiveDataCallback ...
type OnReceiveDataCallback func(symbol string, data *QuoteData)

// OnErrorCallback ...
type OnErrorCallback func(err error, context string)
