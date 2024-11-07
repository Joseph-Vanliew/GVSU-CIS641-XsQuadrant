package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Meeting struct {
	gorm.Model
	MeetingId uuid.UUID `gorm:"unique"`
	AdminID   string    `gorm:"foreignKey: Email"`
	PeerId    uuid.UUID `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
