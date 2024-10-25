package gormDatabase

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string  `gorm:"not null;unique"`
	Age     int64   `gorm:"not null"`
	Profile Profile `gorm:"constraint:OnDelete:CASCADE;"`
}
type Profile struct {
	gorm.Model
	UserID            uint   `gorm:"not null"`
	Bio               string `json:"bio"`
	ProfilePictureURL string `json:"profile_picture_url"`
}
