package client



type StonfiClientOptions struct {
	BaseURL string
}

type MarketListResponse struct {
	Markets []Market `json:"markets"`
}

type Market struct {
	Pair     string `json:"pair"`
	Volume   string `json:"volume"`
	Price    string `json:"price"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

