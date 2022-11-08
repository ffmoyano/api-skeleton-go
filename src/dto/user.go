package dto

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Roles     []Role    `json:"roles"`
	CreatedAt time.Time `json:"createdAt"`
}
