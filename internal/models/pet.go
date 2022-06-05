package models

import (
	"time"
)

type Pet struct {
	ID       uint64 `json:"-"`
	UUID     string `json:"uuid,omitempty"`
	OwnerID  uint64 `json:"owner_id,omitempty" db:"owner_id"`
	PlanUUID string `json:"plan_uuid,omitempty"`
	Name     string `json:"name,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
