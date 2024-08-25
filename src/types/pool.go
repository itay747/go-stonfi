package types

type PoolInfoResponse struct {
	PoolList []struct {
		Address                    string `json:"address"`
		RouterAddress              string `json:"router_address"`
		Reserve0                   string `json:"reserve0"`
		Reserve1                   string `json:"reserve1"`
		Token0Address              string `json:"token0_address"`
		Token1Address              string `json:"token1_address"`
		LpTotalSupply              string `json:"lp_total_supply"`
		LpTotalSupplyUsd           string `json:"lp_total_supply_usd"`
		LpFee                      string `json:"lp_fee"`
		ProtocolFee                string `json:"protocol_fee"`
		RefFee                     string `json:"ref_fee"`
		ProtocolFeeAddress         string `json:"protocol_fee_address"`
		CollectedToken0ProtocolFee string `json:"collected_token0_protocol_fee"`
		CollectedToken1ProtocolFee string `json:"collected_token1_protocol_fee"`
		LpPriceUsd                 string `json:"lp_price_usd"`
		Apy1D                      string `json:"apy_1d,omitempty"`
		Apy7D                      string `json:"apy_7d,omitempty"`
		Apy30D                     string `json:"apy_30d,omitempty"`
		Deprecated                 bool   `json:"deprecated"`
	} `json:"pool_list"`
}
