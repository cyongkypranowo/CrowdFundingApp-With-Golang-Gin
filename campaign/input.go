package campaign

import "crowdfunding/user"

type GetCampaignDetailInput struct {
	ID uint64 `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

type CreateCampaignImageInput struct {
	CampaignID uint64 `form:"campaign_id" binding:"required"`
	// FileName   string `form:"file_name" binding:"required"`
	IsPrimary bool `form:"is_primary"`
	User      user.User
}
