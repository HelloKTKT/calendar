package main

import (
	"github.com/gin-gonic/gin"
	. "time-manager/handlers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/register", AddUser)
	router.POST("/login", Login)

	router.POST("/calendar-reminder", AddCalendarReminder)
	router.DELETE("/calendar-reminder/:id", DeleteCalendarReminder)
	router.PUT("/calendar-reminder/:id", UpdateCalendarReminder)
	router.GET("/calendar-reminder", ListCalendarReminder)
	return router
}
