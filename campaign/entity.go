package campaign

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	ID               uint64
	UserID           uint64
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	Perks            string
	BackerCount      int
	Slug             string
	CreatedBy        uint64
	UpdatedBy        uint64
	DeletedBy        uint64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         uint64
	CampaignID uint64
	FileName   string
	IsPrimary  int
	CreatedBy  uint64
	UpdatedBy  uint64
	DeletedBy  uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
