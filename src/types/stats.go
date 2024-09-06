package types

import "time"

// Aggregated stats for a pool.
type DexStatsResponse struct {
	Since time.Time `json:"since"`
	Until time.Time `json:"until"`
	Stats struct {
		Tvl           string `json:"tvl"`            // Total Value Locked
		VolumeUsd     string `json:"volume_usd"`     // Volume in USD
		Trades        int    `json:"trades"`         // Number of trades
		UniqueWallets int    `json:"unique_wallets"` // Number of unique wallets
	} `json:"stats"`
}

type Operation struct {
	ProtocolFeeAmount        string `json:"protocol_fee_amount"`
	FeeAssetAddress          string `json:"fee_asset_address"`
	RouterAddress            string `json:"router_address"`
	Asset0Reserve            string `json:"asset0_reserve"`
	PoolTxTimestamp          string `json:"pool_tx_timestamp"`
	DestinationWalletAddress string `json:"destination_wallet_address"`
	OperationType            string `json:"operation_type"`
	WalletTxHash             string `json:"wallet_tx_hash"`
	ExitCode                 string `json:"exit_code"`
	Asset0Address            string `json:"asset0_address"`
	Asset0Amount             string `json:"asset0_amount"`
	Asset0Delta              string `json:"asset0_delta"`
	WalletTxTimestamp        string `json:"wallet_tx_timestamp"`
	WalletTxLt               string `json:"wallet_tx_lt"`
	LpTokenSupply            string `json:"lp_token_supply"`
	Asset1Delta              string `json:"asset1_delta"`
	Asset1Reserve            string `json:"asset1_reserve"`
	LpTokenDelta             string `json:"lp_token_delta"`
	Asset1Amount             string `json:"asset1_amount"`
	PoolAddress              string `json:"pool_address"`
	LpFeeAmount              string `json:"lp_fee_amount"`
	PoolTxHash               string `json:"pool_tx_hash"`
	ReferralFeeAmount        string `json:"referral_fee_amount"`
	ReferralAddress          string `json:"referral_address"`
	WalletAddress            string `json:"wallet_address"`
	Asset1Address            string `json:"asset1_address"`
	PoolTxLt                 int64  `json:"pool_tx_lt"`
	Success                  bool   `json:"success"`
}
type Operations []struct {
	Operation Operation `json:"operation"`
	Asset0Info Asset `json:"asset0_info"`
	Asset1Info Asset `json:"asset1_info"`
}
type OperationsStatsResponse struct {
	Operations Operations `json:"operations"`		
}

type PoolStatsResponse struct {
	Since time.Time `json:"since"`
	Until time.Time `json:"until"`
	Stats []struct {
		PoolAddress    string `json:"pool_address"`
		RouterAddress  string `json:"router_address"`
		URL            string `json:"url"`
		BaseID         string `json:"base_id"`
		BaseName       string `json:"base_name"`
		BaseSymbol     string `json:"base_symbol"`
		QuoteID        string `json:"quote_id"`
		QuoteName      string `json:"quote_name"`
		QuoteSymbol    string `json:"quote_symbol"`
		LastPrice      string `json:"last_price"`
		BaseVolume     string `json:"base_volume"`
		QuoteVolume    string `json:"quote_volume"`
		BaseLiquidity  string `json:"base_liquidity"`
		QuoteLiquidity string `json:"quote_liquidity"`
		LpPrice        string `json:"lp_price"`
		LpPriceUsd     string `json:"lp_price_usd"`
		Apy            string `json:"apy"`
	} `json:"stats"`
	UniqueWalletsCount int `json:"unique_wallets_count"`
}

