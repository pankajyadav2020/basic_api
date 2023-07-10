package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password []byte `json:"password,omitempty"`
}
