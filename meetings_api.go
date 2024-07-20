package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type MeetingAPI struct {
	db        *gorm.DB
	listCache map[int][]Meeting
}

func NewMeetingAPI(db *gorm.DB) *MeetingAPI {
	return &MeetingAPI{db: db, listCache: make(map[int][]Meeting)}
}

func (m *MeetingAPI) List(context *gin.Context) {
	var meetings []Meeting

	offset, err := strconv.Atoi(context.DefaultQuery("offset", "0"))
	if err != nil {
		offset = 0
	}
	if value, ok := m.listCache[offset]; ok {
		meetings = value
	} else if result := m.db.Order("id desc").Limit(10).Offset(offset).Find(&meetings); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	} else {
		m.listCache[offset] = meetings
	}
	context.JSON(http.StatusOK, meetings)
}

func (m *MeetingAPI) Create(context *gin.Context) {
	var meeting Meeting
	if err := context.ShouldBindJSON(&meeting); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := m.db.Create(&meeting); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	context.JSON(http.StatusOK, meeting)
}
