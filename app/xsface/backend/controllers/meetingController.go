package controllers

import (
	"gorm.io/gorm"
	"xsface/initializers"
)

type MeetingController struct {
	db *gorm.DB
}

func NewMeetingController() *MeetingController {
	return &MeetingController{
		db: initializers.DB, // Use the global DB connection from initializers
	}
}
