package models

import (
	"github.com/google/uuid"
	"time"
)

type Group struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	OwnerId   uuid.UUID `json:"owner_id"`
	MaxUsers  int       `json:"max_users"`
	CreatedAt time.Time `json:"created_at"`
}
