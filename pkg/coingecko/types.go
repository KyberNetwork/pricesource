package coingecko

type PriceResponse struct {
	Usd float64 `json:"usd"`
}

type CoingeckoCoin struct {
	Id        string            `json:"id"`
	Symbol    string            `json:"symbol"`
	Name      string            `json:"name"`
	Platforms map[string]string `json:"platforms"`
}
