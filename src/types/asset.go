package types

type AssetKind string

const (
	AssetKindTon    AssetKind = "Ton"
	AssetKindWton   AssetKind = "Wton"
	AssetKindJetton AssetKind = "Jetton"
)

type Asset struct {
	ContractAddress    string    `json:"contract_address"`
	Symbol             string    `json:"symbol"`
	DisplayName        string    `json:"display_name"`
	DexPriceUsd        string    `json:"dex_price_usd,omitempty"`
	ImageURL           string    `json:"image_url"`
	DexUsdPrice        string    `json:"dex_usd_price,omitempty"`
	Kind               AssetKind `json:"kind"`
	ThirdPartyPriceUsd string    `json:"third_party_price_usd,omitempty"`
	ThirdPartyUsdPrice string    `json:"third_party_usd_price,omitempty"`
	Tags               []string  `json:"tags"`
	Decimals           int       `json:"decimals"`
	Priority           int       `json:"priority"`
	DefaultSymbol      bool      `json:"default_symbol"`
	Taxable            bool      `json:"taxable"`
	Blacklisted        bool      `json:"blacklisted"`
	Community          bool      `json:"community"`
	Deprecated         bool      `json:"deprecated"`
}

// Struct for `/v1/assets/search` and `/v1/assets/{addr_str}`
type AssetResponse struct {
	Asset Asset `json:"asset"`
}

// Struct for `/v1/assets/` 
type AssetListResponse struct {
	AssetList []Asset `json:"asset_list"`
}



type AssetList struct {
	ContractAddress string `json:"contract_address"`
	Kind            string `json:"kind"`
	DexPriceUsd     string `json:"dex_price_usd"`
	WalletAddress   string `json:"wallet_address"`
	Balance         string `json:"balance"`
	Meta            struct {
		Symbol      string `json:"symbol"`
		DisplayName string `json:"display_name"`
		ImageURL    string `json:"image_url"`
		Decimals    int    `json:"decimals"`
	} `json:"meta"`
	Tags []string `json:"tags"`
}
type SearchAssetsResponse struct {
	AssetList AssetList  `json:"asset_list"`
}
type QueryWalletBallanceResponse struct {
	AssetList AssetList `json:"asset_list"`
}