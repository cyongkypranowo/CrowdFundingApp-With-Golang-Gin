package campaign

import (
	"crowdfunding/user"
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	ID               uint64
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
	UserID           uint64 `gorm:"foreignkey:UserID"`
	User             user.User
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
