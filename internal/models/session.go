package models

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserId    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Expires   time.Time `json:"expires" gorm:"not null"`
	Data      string    `json:"data" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
}
