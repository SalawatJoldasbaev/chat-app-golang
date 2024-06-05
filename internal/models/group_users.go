package models

import (
	"github.com/google/uuid"
	"time"
)

type GroupUser struct {
	GroupId   uuid.UUID `json:"group_id"`
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
