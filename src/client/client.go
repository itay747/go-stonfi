package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/imroc/req/v3"
	"github.com/itay747/go-stonfi/src/types"
)

var BaseURLStr = "https://api.ston.fi/v1"

type StonfiClient struct {
	Client *req.Client
}

// NewStonfiClient creates a new API client for the Ston.fi service.
func NewStonfiClient() *StonfiClient {
	client := req.C().
		SetBaseURL(BaseURLStr).
		SetCommonContentType("application/json").
		SetCommonHeaders(map[string]string{
			"Accept-Encoding": "gzip, deflate, br",
			"Accept":          "application/json",
			"Content-Type":    "application/json",
			"User-Agent":      "go-stonfi/v0.1.0",
		})

	return &StonfiClient{
		Client: client, // You forgot to assign the created client to the StonfiClient struct.
	}
}

func (c *StonfiClient) buildQueryParams(base string, params url.Values) string {
	u, _ := url.Parse(c.Client.BaseURL + base)
	q := u.Query()
	for key, values := range params {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *StonfiClient) request(ctx context.Context, method, url string, body interface{}, response interface{}) error {
	req := c.Client.R().SetContext(ctx).SetBody(body)
	resp, err := req.Send(method, url)
	if err != nil {
		return fmt.Errorf("request error: %s, status code: %d", err, resp.StatusCode)
	}
	if !resp.IsSuccessState() {
		return fmt.Errorf("API error: %s, status code: %d", resp.String(), resp.StatusCode)
	}
	if err = resp.UnmarshalJson(response); err != nil {
		return fmt.Errorf("JSON unmarshal error: %s", err)
	}
	return nil
}

func (c *StonfiClient) GetAsset(ctx context.Context, assetAddress string) (*types.AssetResponse, error) {
	endpoint := fmt.Sprintf("/v1/asset/%s", assetAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response types.AssetResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return &response, err
	}
	return &response, nil
}

// GetAssets fetches details for all assets.
func (c *StonfiClient) GetAssets(ctx context.Context) (*types.AssetListResponse, error) {
	endpoint := "/assets"
	url := c.buildQueryParams(endpoint, nil)
	var response types.AssetListResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetFarm fetches details for a single farm.
func (c *StonfiClient) GetFarm(ctx context.Context, farmAddress string) (*types.FarmResponse, error) {
	endpoint := fmt.Sprintf("/farm/%s", farmAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response types.FarmResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetFarms fetches details for all farms.
func (c *StonfiClient) GetFarms(ctx context.Context) (*types.FarmListResponse, error) {
	url := c.buildQueryParams("/farms", nil)
	var response types.FarmListResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPools fetches details for all pools.
func (c *StonfiClient) GetPools(ctx context.Context) (*types.PoolListResponse, error) {
	url := c.buildQueryParams("/v1/pools", nil)
	var response types.PoolListResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPool fetches details for a single pool.
func (c *StonfiClient) GetPool(ctx context.Context, poolAddress string) (*types.PoolResponse, error) {
	endpoint := fmt.Sprintf("/v1/pool/%s", poolAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response types.PoolResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetSwapStatus fetches the status of a specific swap operation.
func (c *StonfiClient) GetSwapStatus(ctx context.Context, routerAddress, ownerAddress, queryId string) (*types.SwapResponse, error) {
	queryParams := url.Values{
		"routerAddress": []string{routerAddress},
		"ownerAddress":  []string{ownerAddress},
		"queryId":       []string{queryId},
	}
	url := c.buildQueryParams("/v1/swap/status", queryParams)
	var response *types.SwapResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetPoolByAddress fetches details of a pool by its address.
func (c *StonfiClient) GetPoolByAddress(ctx context.Context, poolAddress string) (*types.PoolListResponse, error) {
	endpoint := fmt.Sprintf("/v1/pools/%s", poolAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response *types.PoolListResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// SimulateSwap performs a simulation of a direct swap between two assets.
func (c *StonfiClient) SimulateSwap(ctx context.Context, offerAddress, askAddress, units, slippageTolerance string) (*types.SwapSimulationResponse, error) {
	queryParams := url.Values{
		"offer_address":      []string{offerAddress},
		"ask_address":        []string{askAddress},
		"units":              []string{units},
		"slippage_tolerance": []string{slippageTolerance},
	}
	url := c.buildQueryParams("/v1/swap/simulate", queryParams)
	var response *types.SwapSimulationResponse
	if err := c.request(ctx, http.MethodPost, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// SimulateReverseSwap performs a simulation of a reverse swap between two assets.
func (c *StonfiClient) SimulateReverseSwap(ctx context.Context, offerAddress, askAddress, units, slippageTolerance string) (*types.SwapSimulationResponse, error) {
	queryParams := url.Values{
		"offer_address":      []string{offerAddress},
		"ask_address":        []string{askAddress},
		"units":              []string{units},
		"slippage_tolerance": []string{slippageTolerance},
	}
	url := c.buildQueryParams("/v1/reverse_swap/simulate", queryParams)
	var response *types.SwapSimulationResponse
	if err := c.request(ctx, http.MethodPost, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *StonfiClient) SearchAssets(ctx context.Context, searchString string, condition *string, walletAddress *string) (*types.SearchAssetsResponse, error) {
	queryParams := url.Values{
		"search_string": []string{searchString},
	}
	if condition != nil {
		queryParams.Add("condition", *condition)
	}
	if walletAddress != nil {
		queryParams.Add("wallet_address", *walletAddress)
	}

	url := c.buildQueryParams("/v1/assets/search", queryParams)

	var response *types.SearchAssetsResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *StonfiClient) QueryAssets(ctx context.Context, condition string, unconditionalAssets []string, walletAddress string) (*types.QueryWalletBallanceResponse, error) {
	queryParams := url.Values{}

	if condition != "" {
		queryParams.Add("condition", condition)
	}

	for _, asset := range unconditionalAssets {
		queryParams.Add("unconditional_assets", asset)
	}

	if walletAddress != "" {
		queryParams.Add("wallet_address", walletAddress)
	}

	url := c.buildQueryParams("/v1/assets/query", queryParams)

	var response *types.QueryWalletBallanceResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
func (c *StonfiClient) GetFarmsByPool(ctx context.Context, poolAddress string) (*types.FarmListResponse, error) {
	endpoint := fmt.Sprintf("/farms_by_pool/%s", poolAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response *types.FarmListResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetWalletAssets fetches details of all assets associated with a specific wallet.
func (c *StonfiClient) GetWalletAssets(ctx context.Context, walletAddress string) (*types.SearchAssetsResponse, error) {
	endpoint := fmt.Sprintf("/wallets/%s/assets", walletAddress)
	url := c.buildQueryParams(endpoint, nil)

	var response *types.SearchAssetsResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
func (c *StonfiClient) GetWalletAsset(ctx context.Context, walletAddressStr string, assetStr string) (*types.WalletAssetsListResponse, error) {
	endpoint := fmt.Sprintf("/v1/wallets/%s/assets/%s", walletAddressStr, assetStr)
	url := c.buildQueryParams(endpoint, nil)
	var response *types.WalletAssetsListResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Ston.fi stats endpoints

// Wallet operations in pools
