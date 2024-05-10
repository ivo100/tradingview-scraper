package tradingview

// SocketInterface ...
type SocketInterface interface {
	AddSymbol(symbol string) error
	RemoveSymbol(symbol string) error
	Init() error
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
	Price  *float64 `mapstructure:"lp"`
	Volume *float64 `mapstructure:"volume"`
	Bid    *float64 `mapstructure:"bid"`
	Ask    *float64 `mapstructure:"ask"`
	Change *float64 `mapstructure:"ch"`
	Time   *int64   `mapstructure:"lp_time"`
}

// Flags ...
type Flags struct {
	Flags []string `json:"flags"`
}

// OnReceiveDataCallback ...
type OnReceiveDataCallback func(symbol string, data *QuoteData)

// OnErrorCallback ...
type OnErrorCallback func(err error, context string)
