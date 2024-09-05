package types

import "time"

// Aggregated stats for a pool.
type PoolAggregatedStatsResponse struct {
	Since time.Time `json:"since"`
	Until time.Time `json:"until"`
	Stats struct {
		Tvl           string `json:"tvl"`
		VolumeUsd     string `json:"volume_usd"`
		Trades        int    `json:"trades"`
		UniqueWallets int    `json:"unique_wallets"`
	} `json:"stats"`
}

