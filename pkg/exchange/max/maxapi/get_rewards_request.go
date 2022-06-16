package max

import "github.com/c9s/requestgen"

//go:generate -command GetRequest requestgen -method GET
//go:generate -command PostRequest requestgen -method POST

func (s *RewardService) NewGetRewardsRequest() *GetRewardsRequest {
	return &GetRewardsRequest{client: s.client}
}

//go:generate GetRequest -url "v2/rewards" -type GetRewardsRequest -responseType []Reward
type GetRewardsRequest struct {
	client requestgen.AuthenticatedAPIClient

	currency *string `param:"currency"`

	// From Unix-timestamp
	from *int64 `param:"from"`

	// To Unix-timestamp
	to *int64 `param:"to"`

	page   *int64 `param:"page"`
	limit  *int64 `param:"limit"`
	offset *int64 `param:"offset"`
}
