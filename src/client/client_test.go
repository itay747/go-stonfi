package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func newTestClient() (*StonfiClient, *http.Client) {
	client := NewStonfiClient(StonfiClientOptions{BaseURL: "https://api.ston.fi"})
	httpClient := client.Client.GetClient()
	httpmock.ActivateNonDefault(httpClient)
	return client, httpClient
}
func TestGetAsset(t *testing.T) {
	client, httpClient := newTestClient()
	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()

	mockResponse := `{
		"asset": {
			"contract_address": "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c",
			"kind": "Jetton",
			"meta": {
				"symbol": "TON",
				"display_name": "TON",
				"image_url": "https://asset.ston.fi/img/EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c/4ecd4687e0b5b8ff21a7fbe03f9d281c26a2dc13eac7b7d16048cc693fe0ec39",
				"decimals": 9
			},
			"tags": ["default_symbol"],
			"dex_price_usd": "6.730000000000000"
		}
	}`

	httpmock.RegisterResponder("GET", "https://api.ston.fi/v1/assets/EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c", httpmock.NewStringResponder(http.StatusOK, mockResponse))
	asset, err := client.GetAsset(context.Background(), "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c")
	assert.NoError(t, err)
	assert.NotNil(t, asset)
	assert.Equal(t, "TON", asset.Assets[0].Symbol)
}

func TestGetWalletAssets(t *testing.T) {
	client, httpClient := newTestClient()
	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()

	mockResponse := `{
		"asset_list": [
			{
				"contract_address": "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c",
				"kind": "Jetton",
				"meta": {
					"symbol": "TON",
					"display_name": "TON",
					"image_url": "https://asset.ston.fi/img/EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c/4ecd4687e0b5b8ff21a7fbe03f9d281c26a2dc13eac7b7d16048cc693fe0ec39",
					"decimals": 9
				},
				"tags": ["default_symbol"],
				"dex_price_usd": "6.730000000000000"
			}
		]
	}`

	httpmock.RegisterResponder("GET", "https://api.ston.fi/v1/wallets/UQ.../assets", httpmock.NewStringResponder(http.StatusOK, mockResponse))

	walletAssets, err := client.GetWalletAssets(context.Background(), "UQ...")
	assert.NoError(t, err)
	assert.NotNil(t, walletAssets)
	assert.Equal(t, "TON", walletAssets.Assets[0].Symbol)
}

func TestGetPools(t *testing.T) {
	client, httpClient := newTestClient()
	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()

	mockResponse := `{
		"pool_list": [
			{
				"address": "EQDAJK0GZ2ZhJlHcuNFrGDSyg70s0OZPFCiy29CNkMRomGFZ",
				"router_address": "EQB3ncyBUTjZUA5EnFKR5_EnOMI9V1tTEAAPaiU71gc4TiUt",
				"reserve0": "703912430",
				"reserve1": "1026",
				"token0_address": "EQAqrTGCkKAeQqe8ivX2fsP7dkPTBWV3_59VHHibeFeS8nPD",
				"token1_address": "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c",
				"lp_total_supply": "849",
				"lp_total_supply_usd": "0.000006648479999487",
				"lp_fee": "20",
				"protocol_fee": "10",
				"ref_fee": "10",
				"protocol_fee_address": "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c",
				"collected_token0_protocol_fee": "7663502681622",
				"collected_token1_protocol_fee": "4968692",
				"lp_price_usd": "7.830954063",
				"apy_1d": "0",
				"apy_7d": "0",
				"apy_30d": "0",
				"deprecated": false
			},
		]
	}`
	httpmock.RegisterResponder("GET", "https://api.ston.fi/v1/pools", httpmock.NewStringResponder(http.StatusOK, mockResponse))

	pools, err := client.GetPools(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, pools)
	assert.Equal(t, "TON Pool", pools[0].Name)
}

func TestGetPool(t *testing.T) {
	client, httpClient := newTestClient()
	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()

	mockResponse := `{
		"pool": {
			"address": "EQDAJK0GZ2ZhJlHcuNFrGDSyg70s0OZPFCiy29CNkMRomGFZ",
			"router_address": "EQB3ncyBUTjZUA5EnFKR5_EnOMI9V1tTEAAPaiU71gc4TiUt",
			"reserve0": "703912430",
			"reserve1": "1026",
			"token0_address": "EQAqrTGCkKAeQqe8ivX2fsP7dkPTBWV3_59VHHibeFeS8nPD",
			"token1_address": "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c",
			"lp_total_supply": "849",
			"lp_total_supply_usd": "0.000006648479999487",
			"lp_fee": "20",
			"protocol_fee": "10",
			"ref_fee": "10",
			"protocol_fee_address": "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c",
			"collected_token0_protocol_fee": "7663502681622",
			"collected_token1_protocol_fee": "4968692",
			"lp_price_usd": "7.830954063",
			"apy_1d": "0",
			"apy_7d": "0",
			"apy_30d": "0",
			"deprecated": false
		}
	}`

	httpmock.RegisterResponder("GET", "https://api.ston.fi/v1/pools/EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c", httpmock.NewStringResponder(http.StatusOK, mockResponse))
	pool, err := client.GetPool(context.Background(), "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c")
	assert.NoError(t, err)
	assert.NotNil(t, pool)
	assert.Equal(t, "TON Pool", pool.Name)
}

func TestGetFarms(t *testing.T) {
	client, httpClient := newTestClient()
	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()

	mockResponse := `{
		"farm_list": [
			{
				"id": "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c",
				"name": "TON Farm",
				"status": "active"
			}
		]
	}`

	httpmock.RegisterResponder("GET", "https://api.ston.fi/v1/farms", httpmock.NewStringResponder(http.StatusOK, mockResponse))
	farms, err := client.GetFarms(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, farms)
	assert.Equal(t, "TON Farm", farms[0].Name)
	assert.Equal(t, "active", farms[0].Status)
}

func TestGetFarm(t *testing.T) {
	client, httpClient := newTestClient()
	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()

	mockResponse := `{
		"farm": {
			"id": "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c",
			"name": "TON Farm",
			"status": "active"
		}
	}`

	httpmock.RegisterResponder("GET", "https://api.ston.fi/v1/farms/EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c", httpmock.NewStringResponder(http.StatusOK, mockResponse))

	farm, err := client.GetFarm(context.Background(), "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c")
	assert.NoError(t, err)
	assert.NotNil(t, farm)
	assert.Equal(t, "TON Farm", farm.Name)
	assert.Equal(t, "activope", farm.Status)
}

func TestGetSimulatedSwapStatus(t *testing.T) {
	client, httpClient := newTestClient()
	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()

	mockResponse := `{
		"status": "completed" 
	}`
	// Assuming the Swagger API's expected output is "completed"
	httpmock.RegisterResponder("GET", "https://api.ston.fi/v1/swap/status?routerAddress=routerAddress&ownerAddress=ownerAddress&queryId=queryId", httpmock.NewStringResponder(http.StatusOK, mockResponse))
	status, err := client.GetSwapStatus(context.Background(), "routerAddress", "ownerAddress", "queryId")
	assert.NoError(t, err)
	assert.NotNil(t, status)
	assert.Equal(t, "completed", status)
}

func TestGetSwapRate(t *testing.T) {
	client, httpClient := newTestClient()
	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()

	// Sample response based on the Swagger API's expected output
	mockResponse := `{
		"swap_rate": "0.151656666",
		"min_ask_units": "45451"
	}`

	url := "https://api.ston.fi/v1/swap/simulate?offer_address=EQBynBO23ywHy_CgarY9NK9FTz0yDsG82PtcbSTQgGoXwiuA&ask_address=EQCM3B12QK1e4yZSf8GtBRT0aLMNyEsBc_DhVfRRtOEffLez&units=300&slippage_tolerance=0.001"

	httpmock.RegisterResponder("POST", url, httpmock.NewStringResponder(http.StatusOK, mockResponse))

	simulation, err := client.SimulateSwap(context.Background(), "EQBynBO23ywHy_CgarY9NK9FTz0yDsG82PtcbSTQgGoXwiuA", "EQCM3B12QK1e4yZSf8GtBRT0aLMNyEsBc_DhVfRRtOEffLez", "300", "0.001")
	assert.NoError(t, err)
	assert.NotNil(t, simulation)

	assert.Equal(t, "0.151656666", simulation.SwapRate)
	assert.Equal(t, "45451", simulation.MinAskUnits)
}
