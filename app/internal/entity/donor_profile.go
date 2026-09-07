package entity

import "time"

type DonorProfile struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	UserID           uint       `gorm:"uniqueIndex;not null" json:"user_id"`
	BloodType        string     `gorm:"not null" json:"blood_type"`
	City             string     `gorm:"not null" json:"city"`
	Latitude         float64    `json:"latitude"`
	Longitude        float64    `json:"longitude"`
	IsAvailable      bool       `gorm:"default:true;not null" json:"is_available"`
	LastDonationDate *time.Time `json:"last_donation_date,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`

	User *User `gorm:"foreignKey:UserID;" json:"user,omitempty"`
}
