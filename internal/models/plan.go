package models

import (
	"time"
)

type Plan struct {
	ID      uint64 `json:"-"`
	UUID    string `json:"uuid,omitempty"`
	OwnerID uint64 `json:"owner_id,omitempty"`
	Name    string `json:"name,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
