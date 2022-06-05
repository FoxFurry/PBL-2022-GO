package models

import (
	"time"
)

type Device struct {
	ID       uint64 `json:"-"`
	UUID     string `json:"uuid,omitempty"`
	OwnerID  uint64 `json:"-"`
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Address  string `json:"address,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
