package types

// Struct for unmarshalling response from `/v1/wallets/{addr_str}/assets`
type WalletAssetsResponse struct {
	Assets []Asset `json:"assets"`
}

// Struct for unmarshalling response from `/v1/wallets/{addr_str}/assets/{asset_address}`
type WalletAssetResponse struct {
	Asset Asset `json:"asset"`
}

// Struct for unmarshalling response from `/v1/wallets/{addr_str}/farms`
type WalletFarmsResponse struct {
	Farms []Farm `json:"farms"`
}

// Struct for unmarshalling response from `/v1/wallets/{addr_str}/farms/{farm_address}`
type WalletFarmResponse struct {
	Farm Farm `json:"farm"`
}

// Struct for unmarshalling response from `/v1/wallets/{addr_str}/pools`
type WalletPoolsResponse struct {
	Pools []Pool `json:"pools"`
}

// Struct for unmarshalling response from `/v1/wallets/{addr_str}/pools/{pool_address}`
type WalletPoolResponse struct {
	Pool Pool `json:"pool"`
}

// Struct for unmarshalling response from `/v1/wallets/{addr_str}/operations`
type WalletOperationsResponse struct {
	Operations Operations `json:"operations"`
}