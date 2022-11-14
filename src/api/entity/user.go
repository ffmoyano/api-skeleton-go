package entity

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Roles     []Role    `json:"roles,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Verified  bool      `json:"-"`
}
