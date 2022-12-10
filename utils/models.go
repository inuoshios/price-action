package utils

type CoinDetails struct {
	Name                     string  `json:"name"`
	Symbol                   string  `json:"symbol"`
	Price                    float64 `json:"current_price"`
	PriceChangePercentage24h float64 `json:"price_change_percentage_24h"`
	MarketCapChange24h       float64 `json:"market_cap_change_24h"`
	ATH                      float64 `json:"ath"`
}

type NftDetails struct {
	Slug     string  `json:"slug"`
	Volume7D float64 `json:"seven_day_volume"`
	Stats    struct {
		Sales7D       int64   `json:"seven_day_sales"`
		VolumeAllTime float64 `json:"total_volume"`
		TotalSupply   int64   `json:"total_supply"`
		MarketCap     float64 `json:"market_cap"`
		FloorPrice    float64 `json:"floor_price"`
	} `json:"stats"`
}

type StockDetails struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency           string  `json:"currency"`
				Symbol             string  `json:"symbol"`
				Timezone           string  `json:"timezone"`
				RegularMarketPrice float64 `json:"regularMarketPrice"`
				ChartPreviousClose float64 `json:"chartPreviousClose"`
				PreviousClose      float64 `json:"previousClose"`
			} `json:"meta"`
		} `json:"result"`
	} `json:"chart"`
}
