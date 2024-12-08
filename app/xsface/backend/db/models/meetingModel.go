package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Meeting struct {
	gorm.Model
	MeetingId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	//CREATE EXTENSION IF NOT EXISTS "uuid-ossp";<-- make sure this is in the schema for our postgres db
	AdminID      string     `gorm:"foreignKey:Email"`
	PeerID       uuid.UUID  `gorm:"type:uuid;unique"`
	ScheduledFor *time.Time `gorm:"default:null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
