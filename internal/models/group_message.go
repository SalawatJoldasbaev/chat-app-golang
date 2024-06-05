package models

import (
	"github.com/google/uuid"
	"time"
)

type GroupMessage struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	GroupId   uuid.UUID `json:"group_id"`
	UserId    uuid.UUID `json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
