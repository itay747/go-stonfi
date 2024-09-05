package types

type FarmInfoResponse struct {
	FarmList []struct {
		MinterAddress      string        `json:"minter_address"`
		PoolAddress        string        `json:"pool_address"`
		RewardTokenAddress string        `json:"reward_token_address"`
		Status             string        `json:"status"`
		MinStakeDurationS  string        `json:"min_stake_duration_s"`
		LockedTotalLp      string        `json:"locked_total_lp"`
		LockedTotalLpUsd   string        `json:"locked_total_lp_usd"`
		Apy                string        `json:"apy"`
		NftInfos           []interface{} `json:"nft_infos"`
		Rewards            []struct {
			Address          string `json:"address"`
			Status           string `json:"status"`
			RemainingRewards string `json:"remaining_rewards"`
			RewardRate24H    string `json:"reward_rate_24h"`
		} `json:"rewards"`
	} `json:"farm_list"`
}

type FarmResponse struct {
	Farm Farm
}

type FarmListResponse struct {
	Farms []Farm `json:"farms"`
}

type Farm struct {
	MinterAddress      string        `json:"minter_address"`
	PoolAddress        string        `json:"pool_address"`
	RewardTokenAddress string        `json:"reward_token_address"`
	Status             string        `json:"status"`
	MinStakeDurationS  string        `json:"min_stake_duration_s"`
	LockedTotalLP      string        `json:"locked_total_lp"`
	LockedTotalLPUSD   string        `json:"locked_total_lp_usd"`
	APY                string        `json:"apy"`
	NftInfos           []interface{} `json:"nft_infos"`
	Rewards            []struct {
		Address          string `json:"address"`
		Status           string `json:"status"`
		RemainingRewards string `json:"remaining_rewards"`
		RewardRate24H    string `json:"reward_rate_24h"`
	} `json:"rewards"`
}
