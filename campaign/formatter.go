package campaign

type CampaignFormatter struct {
	ID               uint64 `json:"id"`
	UserID           uint64 `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
	}
	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	}
	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter = append(campaignFormatter, FormatCampaign(campaign))
	}

	return campaignFormatter
}
