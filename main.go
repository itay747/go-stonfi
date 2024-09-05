package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/itay747/go-stonfi/src/client"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	apiKey := os.Getenv("STONFI_API_KEY") 
	
	if apiKey == "" {
		errorMessage("STONFI_API_KEY environment variable is not set")
		os.Exit(1)
	}

	action := flag.String("action", "", "Action to perform (assets, asset, wallet-assets, swap-simulate, pools, farms, farm-by-pool)")
	assetAddress := flag.String("asset-address", "", "Asset address (for asset, wallet-assets, swap-simulate actions)")
	walletAddress := flag.String("wallet-address", "", "Wallet address (for wallet-assets, wallet-pools, wallet-farms actions)")
	offerAddress := flag.String("offer-address", "", "Offer asset address (for swap-simulate action)")
	askAddress := flag.String("ask-address", "", "Ask asset address (for swap-simulate action)")
	amount := flag.String("amount", "0", "Amount to swap (for swap-simulate action)")
	slippage := flag.String("slippage", "0.01", "Slippage tolerance (for swap-simulate action)")
	poolAddress := flag.String("pool-address", "", "Pool address (for pool and farm-by-pool actions)")

	flag.Parse()

	client := client.NewStonfiClient()

	ctx := context.Background()

	switch strings.ToLower(*action) {
	case "assets":
		handleAssets(ctx, client)
	case "asset":
		validateRequiredFlag("asset-address", *assetAddress)
	case "wallet-assets":
		validateRequiredFlag("wallet-address", *walletAddress)
		handleWalletAssets(ctx, client, *walletAddress)
	case "swap-simulate":
		validateRequiredFlag("offer-address", *offerAddress)
		validateRequiredFlag("ask-address", *askAddress)
		handleSwapSimulate(ctx, client, *offerAddress, *askAddress, *amount, *slippage)
	// case "pools":
	// 	handlePools(ctx, client)
	case "pool":
		validateRequiredFlag("pool-address", *poolAddress)
		handlePoolByAddress(ctx, client, *poolAddress)
	case "wallet-pools":
		validateRequiredFlag("wallet-address", *walletAddress)
		// handleWalletPools(ctx, client, *walletAddress)
	case "farms":
		handleFarms(ctx, client)
	case "farm":
		validateRequiredFlag("asset-address", *assetAddress)
		handleFarm(ctx, client, *assetAddress)
	case "wallet-farms":
		validateRequiredFlag("wallet-address", *walletAddress)
	//	handleWalletFarms(ctx, client, *walletAddress)
	case "farm-by-pool":
		validateRequiredFlag("pool-address", *poolAddress)
	//	handleFarmByPool(ctx, client, *poolAddress)
	default:
		errorMessage("Invalid action. Valid actions are: assets, asset, wallet-assets, swap-simulate, pools, pool, wallet-pools, farms, farm, wallet-farms, farm-by-pool")
		os.Exit(1)
	}
}

func handleAssets(ctx context.Context, client *client.StonfiClient) {
	infoMessage("Fetching all DEX assets...")
	assets, err := client.GetAssets(ctx)
	if err != nil {
		errorMessage(fmt.Sprintf("Error fetching assets: %v", err))
		os.Exit(1)
	}
	successMessage("Assets fetched successfully!")
	printJSON("Assets", assets)
}

func handleWalletAssets(ctx context.Context, client *client.StonfiClient, walletAddress string) {
	infoMessage(fmt.Sprintf("Fetching wallet assets for %s...", walletAddress))
	walletAssets, err := client.GetWalletAssets(ctx, walletAddress)
	if err != nil {
		errorMessage(fmt.Sprintf("Error fetching wallet assets: %v", err))
		os.Exit(1)
	}
	successMessage("Wallet assets fetched successfully!")
	printJSON("Wallet Assets", walletAssets)
}

func handleSwapSimulate(ctx context.Context, client *client.StonfiClient, offerAddress, askAddress, amount, slippage string) {
	infoMessage(fmt.Sprintf("Simulating swap: %s -> %s, amount: %s, slippage: %s", offerAddress, askAddress, amount, slippage))
	swapSimulation, err := client.SimulateSwap(ctx, offerAddress, askAddress, amount, slippage)
	if err != nil {
		errorMessage(fmt.Sprintf("Error simulating swap: %v", err))
		os.Exit(1)
	}
	successMessage("Swap simulation completed successfully!")
	printJSON("Swap Simulation", swapSimulation)
}


func handlePoolByAddress(ctx context.Context, client *client.StonfiClient, poolAddress string) {
	infoMessage(fmt.Sprintf("Fetching pool details for %s...", poolAddress))
	pool, err := client.GetPoolByAddress(ctx, poolAddress)
	if err != nil {
		errorMessage(fmt.Sprintf("Error fetching pool: %v", err))
		os.Exit(1)
	}
	successMessage("Pool details fetched successfully!")
	printJSON("Pool", pool)
}

// func handleWalletPools(ctx context.Context, client *client.StonfiClient, walletAddress string) {
// 	infoMessage(fmt.Sprintf("Fetching wallet pools for %s...", walletAddress))
// 	walletPools, err := client.GetWalletPools(ctx, walletAddress)
// 	if err != nil {
// 		errorMessage(fmt.Sprintf("Error fetching wallet pools: %v", err))
// 		os.Exit(1)
// 	}
// 	successMessage("Wallet pools fetched successfully!")
// 	printJSON("Wallet Pools", walletPools)
// }

func handleFarms(ctx context.Context, client *client.StonfiClient) {
	infoMessage("Fetching all DEX farms...")
	farms, err := client.GetFarms(ctx)
	if err != nil {
		errorMessage(fmt.Sprintf("Error fetching farms: %v", err))
		os.Exit(1)
	}
	successMessage("Farms fetched successfully!")
	printJSON("Farms", farms)
}

func handleFarm(ctx context.Context, client *client.StonfiClient, farmAddress string) {
	infoMessage(fmt.Sprintf("Fetching farm details for %s...", farmAddress))
	farm, err := client.GetFarm(ctx, farmAddress)
	if err != nil {
		errorMessage(fmt.Sprintf("Error fetching farm: %v", err))
		os.Exit(1)
	}
	successMessage("Farm details fetched successfully!")
	printJSON("Farm", farm)
}

// func handleWalletFarms(ctx context.Context, client *client.StonfiClient, walletAddress string) {
// 	infoMessage(fmt.Sprintf("Fetching wallet farms for %s...", walletAddress))
// 	walletFarms, err := client.GetWalletFarms(ctx, walletAddress)
// 	if err != nil {
// 		errorMessage(fmt.Sprintf("Error fetching wallet farms: %v", err))
// 		os.Exit(1)
// 	}
// 	successMessage("Wallet farms fetched successfully!")
// 	printJSON("Wallet Farms", walletFarms)
// }

func handleFarmsByPool(ctx context.Context, client *client.StonfiClient, poolAddress string) {
	infoMessage(fmt.Sprintf("Fetching farms for pool %s...", poolAddress))
	farm, err := client.GetFarm(ctx, poolAddress)
	if err != nil {
		errorMessage(fmt.Sprintf("Error fetching farms for pool: %v", err))
		os.Exit(1)
	}
	successMessage("Farms for pool fetched successfully!")
	fmt.Printf("Farm is: %+v", farm)
}

func printJSON(title string, data interface{}) {
	fmt.Printf("\n%s:\n", color.New(color.FgCyan, color.Bold).SprintFunc()(title))
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		errorMessage(fmt.Sprintf("Error marshalling JSON: %v", err))
		return
	}
	fmt.Println(string(output))
}

func validateRequiredFlag(flagName, flagValue string) {
	if flagValue == "" {
		errorMessage(fmt.Sprintf("Error: %s is required", flagName))
		os.Exit(1)
	}
}

func infoMessage(message string) {
	color.New(color.FgYellow).Println("\n[INFO]", message)
}

func successMessage(message string) {
	color.New(color.FgGreen).Println("\n[SUCCESS]", message)
}

func errorMessage(message string) {
	color.New(color.FgRed).Println("\n[ERROR]", message)
}
