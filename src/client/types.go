package client

import (
	"time"
)



type StonfiClientOptions struct {
	BaseURL string
}



type DexStatsResponse struct {
	TVL        string `json:"tvl"`
	Volume24h  string `json:"volume_24h"`
	TotalSwaps int    `json:"total_swaps"`
}

type OperationStatsResponse struct {
	Operations []Operation `json:"operations"`
}

type Operation struct {
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Value     string    `json:"value"`
	Status    string    `json:"status"`
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

