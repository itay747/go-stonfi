package types

type AssetKind string

const (
	AssetKindTon    AssetKind = "Ton"
	AssetKindWton   AssetKind = "Wton"
	AssetKindJetton AssetKind = "Jetton"
)

type AssetInfoResponse struct {
	AssetList []struct {
		ContractAddress    string   `json:"contract_address"`
		Symbol             string   `json:"symbol"`
		DisplayName        string   `json:"display_name"`
		DexPriceUsd        string   `json:"dex_price_usd,omitempty"`
		ImageURL           string   `json:"image_url"`
		DexUsdPrice        string   `json:"dex_usd_price,omitempty"`
		Kind               string   `json:"kind"`
		ThirdPartyPriceUsd string   `json:"third_party_price_usd,omitempty"`
		ThirdPartyUsdPrice string   `json:"third_party_usd_price,omitempty"`
		Tags               []string `json:"tags"`
		Decimals           int      `json:"decimals"`
		Priority           int      `json:"priority"`
		DefaultSymbol      bool     `json:"default_symbol"`
		Taxable            bool     `json:"taxable"`
		Blacklisted        bool     `json:"blacklisted"`
		Community          bool     `json:"community"`
		Deprecated         bool     `json:"deprecated"`
	} `json:"asset_list"`
}
