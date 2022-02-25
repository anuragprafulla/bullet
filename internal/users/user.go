package users

import (
	"time"
)

type User struct {
	ID          string    `gorm:"primary_key" json:"id,omitempty"`
	FirstName   string    `json:"firstname,omitempty"`
	LastName    string    `json:"lastname,omitempty"`
	EmailId     string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phone,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
