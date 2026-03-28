package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirebaseUID   string `gorm:"uniqueIndex;size:128;not null" json:"firebase_uid"`
	Email         string `gorm:"uniqueIndex;size:255;not null" json:"email"`
}