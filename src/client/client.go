package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/imroc/req/v3"
	"github.com/itay747/go-stonfi/src/types"
)

var BaseURLStr = "https://api.ston.fi/v1"
var EarliestDate = time.Date(2022, 11, 17, 0, 0, 0, 0, time.UTC)

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

// Check if the time range is valid for retrieving historical data from `/v1/`
func checkValidTimeRange(startDate, endDate time.Time) error {
	if endDate.Before(startDate) || endDate.Equal(startDate) {
		return fmt.Errorf("endDate must be after startDate (endDate: %s, startDate: %s)", endDate, startDate)
	} else if endDate.Sub(startDate) > time.Hour*24 {
		return fmt.Errorf("time range must be less than 24 hours (timeSpan: %s)", endDate.Sub(startDate))
	} else if startDate.Before(EarliestDate) || endDate.Before(EarliestDate) {
		return fmt.Errorf("time range must be after ston.fi mainnet launch date of %s (startDate: %s, endDate: %s)", EarliestDate, startDate, endDate)
	}
	return nil
}

// GetStats retrieves aggregated statistics for a specific time range from `/v1/stats/dex`
func (c *StonfiClient) GetStats(ctx context.Context, startDate, endDate time.Time) (*types.DexStatsResponse, error) {

	if err := checkValidTimeRange(startDate, endDate); err != nil {
		return nil, err
	}

	queryParams := url.Values{
		"start_date": []string{startDate.Format(time.RFC3339)},
		"end_date":   []string{endDate.Format(time.RFC3339)},
	}
	url := c.buildQueryParams("/v1/stats/dex", queryParams)
	var response *types.DexStatsResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieve historical swap data for a specific time range `/v1/stats/operations`
func (c *StonfiClient) GetHistoricalSwaps(ctx context.Context, startDate, endDate time.Time) (*types.OperationsStatsResponse, error) {
	if err := checkValidTimeRange(startDate, endDate); err != nil {
		return nil, err
	}
	queryParams := url.Values{
		"start_date": []string{startDate.Format(time.RFC3339)},
		"end_date":   []string{endDate.Format(time.RFC3339)},
	}
	url := c.buildQueryParams("/v1/stats/operations", queryParams)
	var response *types.OperationsStatsResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieve historical pool data for a specific time range `/v1/stats/pools`
func (c *StonfiClient) GetPoolStats(ctx context.Context, startDate, endDate time.Time) (*types.PoolStatsResponse, error) {
	if err := checkValidTimeRange(startDate, endDate); err != nil {
		return nil, err
	}
	queryParams := url.Values{
		"start_date": []string{startDate.Format(time.RFC3339)},
		"end_date":   []string{endDate.Format(time.RFC3339)},
	}
	url := c.buildQueryParams("/v1/stats/pools", queryParams)
	var response *types.PoolStatsResponse
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

// GetWalletAsset fetches details of a specific asset associated with a specific wallet.
func (c *StonfiClient) GetWalletAsset(ctx context.Context, walletAddressStr string, assetStr string) (*types.WalletAssetResponse, error) {
	endpoint := fmt.Sprintf("/v1/wallets/%s/assets/%s", walletAddressStr, assetStr)
	url := c.buildQueryParams(endpoint, nil)
	var response *types.WalletAssetResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Wallet-specific farms
func (c *StonfiClient) GetWalletFarms(ctx context.Context, walletAddress string) (*types.FarmListResponse, error) {
	endpoint := fmt.Sprintf("/wallets/%s/farms", walletAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response *types.FarmListResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
func (c *StonfiClient) GetWalletFarm(ctx context.Context, walletAddress string, farmAddress string) (*types.FarmResponse, error) {
	endpoint := fmt.Sprintf("/wallets/%s/farms/%s", walletAddress, farmAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response *types.FarmResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Wallet-specific pools
func (c *StonfiClient) GetWalletPools(ctx context.Context, walletAddress string) (*types.PoolListResponse, error) {
	endpoint := fmt.Sprintf("/wallets/%s/pools", walletAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response *types.PoolListResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
func (c *StonfiClient) GetWalletPool(ctx context.Context, walletAddress string, poolAddress string) (*types.PoolResponse, error) {
	endpoint := fmt.Sprintf("/wallets/%s/pools/%s", walletAddress, poolAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response *types.PoolResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Wallet-specific operations
func (c *StonfiClient) GetWalletOperations(ctx context.Context, walletAddress string) (*types.WalletOperationsResponse, error) {
	endpoint := fmt.Sprintf("/wallets/%s/operations", walletAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response *types.WalletOperationsResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
