package client

import "github.com/imroc/req/v3"

// StonfiClient struct represents the client for the Ston.fi API.
type StonfiClient struct {
	Client  *req.Client
	BaseURL string
}

// StonfiClientOptions defines options for initializing a StonfiClient.
type StonfiClientOptions struct {
	BaseURL string
}

type AssetsResponse struct {
	Assets []Asset `json:"asset_list"`
}

type Asset struct {
	ContractAddress    string   `json:"contract_address"`
	Symbol             string   `json:"symbol"`
	DisplayName        string   `json:"display_name"`
	Priority           int      `json:"priority"`
	ImageURL           string   `json:"image_url"`
	Decimals           int      `json:"decimals"`
	Kind               string   `json:"kind"`
	Deprecated         bool     `json:"deprecated"`
	Community          bool     `json:"community"`
	Blacklisted        bool     `json:"blacklisted"`
	DefaultSymbol      bool     `json:"default_symbol"`
	Taxable            bool     `json:"taxable"`
	Tags               []string `json:"tags"`
	ThirdPartyUsdPrice string   `json:"third_party_usd_price,omitempty"`
	ThirdPartyPriceUsd string   `json:"third_party_price_usd,omitempty"`
	DexUsdPrice        string   `json:"dex_usd_price,omitempty"`
	DexPriceUsd        string   `json:"dex_price_usd,omitempty"`
}

type FarmInfoResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type PoolInfoResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SwapStatusResponse struct {
	Status string `json:"status"`
}

type SwapResponse struct {
	OfferAddress      string `json:"offer_address"`
	AskAddress        string `json:"ask_address"`
	OfferJettonWallet string `json:"offer_jetton_wallet"`
	AskJettonWallet   string `json:"ask_jetton_wallet"`
	RouterAddress     string `json:"router_address"`
	PoolAddress       string `json:"pool_address"`
	OfferUnits        string `json:"offer_units"`
	AskUnits          string `json:"ask_units"`
	SlippageTolerance string `json:"slippage_tolerance"`
	MinAskUnits       string `json:"min_ask_units"`
	SwapRate          string `json:"swap_rate"`
	PriceImpact       string `json:"price_impact"`
	FeeAddress        string `json:"fee_address"`
	FeeUnits          string `json:"fee_units"`
	FeePercent        string `json:"fee_percent"`
}
