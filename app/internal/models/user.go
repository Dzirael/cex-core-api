package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID  `json:"user_id"`
	Email     string     `json:"email"`
	Name      string     `json:"name"`
	Surname   string     `json:"surname"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
