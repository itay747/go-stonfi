package types

// SwapSimulationResponse represents the response structure for a swap simulation.
type SwapSimulationResponse struct {
	AskAddress        string `json:"ask_address"`
	AskJettonWallet   string `json:"ask_jetton_wallet"`
	AskUnits          string `json:"ask_units"`
	FeeAddress        string `json:"fee_address"`
	FeePercent        string `json:"fee_percent"`
	FeeUnits          string `json:"fee_units"`
	MinAskUnits       string `json:"min_ask_units"`
	OfferAddress      string `json:"offer_address"`
	OfferJettonWallet string `json:"offer_jetton_wallet"`
	OfferUnits        string `json:"offer_units"`
	PoolAddress       string `json:"pool_address"`
	PriceImpact       string `json:"price_impact"`
	RouterAddress     string `json:"router_address"`
	SlippageTolerance string `json:"slippage_tolerance"`
	SwapRate          string `json:"swap_rate"`
}
type SwapResponse struct {
	TransactionID  string `json:"transaction_id"`
	Status         string `json:"status"`
	PriceImpact    string `json:"price_impact"`
	AmountIn       string `json:"amount_in"`
	AmountOut      string `json:"amount_out"`
	TokenIn        string `json:"token_in"`
	TokenOut       string `json:"token_out"`
}

type SwapStatusResponse struct {
	Status        string `json:"status"`
	Details       string `json:"details"`
}
