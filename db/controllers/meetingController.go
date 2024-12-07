package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"time"
	"v/db/initializers"
	"v/db/models"
)

type MeetingController struct {
	db *gorm.DB
}

func NewMeetingController() *MeetingController {
	return &MeetingController{db: initializers.DB}
}

// GetMeeting retrieves a specific meeting by ID
func (c *MeetingController) GetMeeting(ctx *gin.Context) {
	id := ctx.Param("id")
	var meeting models.Meeting

	if err := c.db.First(&meeting, "meeting_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Meeting not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch meeting"})
		}
		return
	}

	ctx.JSON(http.StatusOK, meeting)
}

// UpdateMeeting updates an existing meeting by ID
func (c *MeetingController) UpdateMeeting(ctx *gin.Context) {
	id := ctx.Param("id")
	var meeting models.Meeting

	// Find meeting by MeetingID
	if err := c.db.First(&meeting, "meeting_id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Meeting not found"})
		return
	}

	var input struct {
		ScheduledFor *time.Time `json:"scheduled_for"` // Allow update of ScheduledFor field
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update meeting fields
	c.db.Model(&meeting).Updates(input)

	ctx.JSON(http.StatusOK, meeting)
}

// DeleteMeeting deletes a specific meeting by ID
func (c *MeetingController) DeleteMeeting(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.db.Delete(&models.Meeting{}, "meeting_id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete meeting"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Meeting deleted"})
}

// ListMeetings retrieves all meetings
func (c *MeetingController) ListMeetings(ctx *gin.Context) {
	var meetings []models.Meeting
	if err := c.db.Find(&meetings).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list meetings"})
		return
	}

	ctx.JSON(http.StatusOK, meetings)
}

// CreateMeeting handles the creation of a new scheduled meeting
func (c *MeetingController) CreateMeeting(ctx *gin.Context) {
	var input struct {
		AdminID      string     `json:"admin_id" binding:"required"`
		PeerID       uuid.UUID  `json:"peer_id" binding:"required"`
		ScheduledFor *time.Time `json:"scheduled_for"` // Optional for instant meetings
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	meeting := models.Meeting{
		MeetingId:    uuid.New(),
		AdminID:      input.AdminID,
		PeerID:       input.PeerID,
		ScheduledFor: input.ScheduledFor,
	}

	if err := c.db.Create(&meeting).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create meeting"})
		return
	}

	ctx.JSON(http.StatusCreated, meeting)
}

// GenerateInstantMeeting creates an instant meeting without storing it in the database
func (c *MeetingController) GenerateInstantMeeting(ctx *gin.Context) {
	// Generate a unique MeetingID and temporary meeting structure
	meeting := models.Meeting{
		MeetingId: uuid.New(),
		AdminID:   "temp-admin-id", // Placeholder for the admin if necessary
		PeerID:    uuid.New(),
		CreatedAt: time.Now(),
	}

	// Generate join URL based on MeetingID
	joinURL := "https://yourapp.com/join/" + meeting.MeetingId.String()

	// Respond with the meeting details and join URL
	ctx.JSON(http.StatusOK, gin.H{"join_url": joinURL, "meeting": meeting})
}
