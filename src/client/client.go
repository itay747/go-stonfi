package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/imroc/req/v3"
)

// NewStonfiClient creates a new API client for the Ston.fi service.
func NewStonfiClient(options StonfiClientOptions) *StonfiClient {
	baseURL := options.BaseURL
	if baseURL == "" {
		baseURL = "https://api.ston.fi"
	}
	if _, err := url.ParseRequestURI(baseURL); err != nil {
		panic(fmt.Errorf("invalid base URL: %w", err))
	}
	client := req.C().
		SetBaseURL(baseURL).
		SetCommonContentType("application/json").
		SetCommonHeaders(map[string]string{
			"Accept-Encoding": "gzip, deflate, br",
			"Accept":          "application/json",
			"Content-Type":    "application/json",
			"User-Agent":      "go-stonfi/1.0.0",
		})

	return &StonfiClient{
		BaseURL: baseURL,
		Client:  client,
	}
}

// buildQueryParams constructs query parameters for a request.
func (c *StonfiClient) buildQueryParams(base string, params url.Values) string {
	u, _ := url.Parse(c.BaseURL + base)
	q := u.Query()
	for key, values := range params {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	u.RawQuery = q.Encode()
	return u.String()
}

// request performs a HTTP request and unmarshals the response into the provided response struct pointer.
func (c *StonfiClient) request(ctx context.Context, method, url string, body interface{}, response interface{}) error {
	req := c.Client.R().SetContext(ctx).SetBody(body)
	resp, err := req.Send(method, url)
	if err != nil {
		return fmt.Errorf("request error: %s, status code: %d", err, resp.StatusCode)
	}
	if err = resp.UnmarshalJson(response); err != nil {
		return fmt.Errorf("JSON unmarshal error: %s", err)
	}
	return nil
}

// GetAssets fetches details for all assets.
func (c *StonfiClient) GetAsset(ctx context.Context, assetAddress string) (AssetsResponse, error) {
	url := fmt.Sprintf("/v1/assets/%s", assetAddress)
	var response AssetsResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return AssetsResponse{}, err
	}
	return response, nil
}

// GetAssets fetches details for all assets.
func (c *StonfiClient) GetAssets(ctx context.Context) (*AssetsResponse, error) {

	url := "/v1/assets"
	var response AssetsResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetFarm fetches details for a single farm.
func (c *StonfiClient) GetFarm(ctx context.Context, farmAddress string) (*FarmInfoResponse, error) {
	endpoint := fmt.Sprintf("/v1/farms/%s", farmAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response FarmInfoResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetFarms fetches details for all farms.
func (c *StonfiClient) GetFarms(ctx context.Context) ([]FarmInfoResponse, error) {
	url := c.buildQueryParams("/v1/farms", nil)
	var response []FarmInfoResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetPools fetches details for all pools.
func (c *StonfiClient) GetPools(ctx context.Context) ([]PoolInfoResponse, error) {
	url := c.buildQueryParams("/v1/pools", nil)
	var response []PoolInfoResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetPool fetches details for a single pool.
func (c *StonfiClient) GetPool(ctx context.Context, poolAddress string) (*PoolInfoResponse, error) {
	endpoint := fmt.Sprintf("/v1/pools/%s", poolAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response PoolInfoResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetSwapStatus fetches the status of a specific swap operation.
func (c *StonfiClient) GetSwapStatus(ctx context.Context, routerAddress, ownerAddress, queryId string) (*SwapStatusResponse, error) {
	queryParams := url.Values{
		"routerAddress": []string{routerAddress},
		"ownerAddress":  []string{ownerAddress},
		"queryId":       []string{queryId},
	}
	url := c.buildQueryParams("/v1/swap/status", queryParams)
	var response SwapStatusResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPoolByAddress fetches details of a pool by its address.
func (c *StonfiClient) GetPoolByAddress(ctx context.Context, poolAddress string) (*PoolInfoResponse, error) {
	endpoint := fmt.Sprintf("/v1/pools/%s", poolAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response PoolInfoResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SimulateSwap performs a simulation of a direct swap between two assets.
func (c *StonfiClient) SimulateSwap(ctx context.Context, offerAddress, askAddress, units, slippageTolerance string) (*SwapResponse, error) {
	queryParams := url.Values{
		"offer_address":      []string{offerAddress},
		"ask_address":        []string{askAddress},
		"units":              []string{units},
		"slippage_tolerance": []string{slippageTolerance},
	}
	url := c.buildQueryParams("/v1/swap/simulate", queryParams)
	var response SwapResponse
	if err := c.request(ctx, http.MethodPost, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SimulateReverseSwap performs a simulation of a reverse swap between two assets.
func (c *StonfiClient) SimulateReverseSwap(ctx context.Context, offerAddress, askAddress, units, slippageTolerance string) (*SwapResponse, error) {
	queryParams := url.Values{
		"offer_address":      []string{offerAddress},
		"ask_address":        []string{askAddress},
		"units":              []string{units},
		"slippage_tolerance": []string{slippageTolerance},
	}
	url := c.buildQueryParams("/v1/reverse_swap/simulate", queryParams)
	var response SwapResponse
	if err := c.request(ctx, http.MethodPost, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetWalletAssets fetches details of all assets associated with a specific wallet.
func (c *StonfiClient) GetWalletAssets(ctx context.Context, walletAddress string) (*AssetsResponse, error) {
	endpoint := fmt.Sprintf("/v1/wallets/%s/assets", walletAddress)
	url := c.buildQueryParams(endpoint, nil)
	var response AssetsResponse
	if err := c.request(ctx, http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
