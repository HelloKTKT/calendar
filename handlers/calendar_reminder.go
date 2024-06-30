package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"strconv"
	"time"
	"time-manager/models"
	"time-manager/utils"
)

type AddParam struct {
	ReminderContent string `form:"reminderContent" binding:"required"`
	ReminderTime    string `form:"reminderTime" binding:"required"`
}

func AddCalendarReminder(context *gin.Context) {

	var param AddParam
	if err := context.ShouldBind(&param); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := getUserIdfromJWT(context)

	layout := "2006-01-02 15:04:05" // Go的time包规定的日期格式字符串
	t, err := time.Parse(layout, param.ReminderTime)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "reminderTime参数校验错误，请输入正确的日期格式（YYYY-MM-dd HH:mm:ss）",
		})
		return
	}

	p := models.CalendarReminder{Content: param.ReminderContent, Time: t, UserId: userId}
	id, err := p.Add()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "add Calendar Reminder fail",
		})
		return
	}
	msg := fmt.Sprintf("insert success %d", id)
	context.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    msg,
	})
}

func DeleteCalendarReminder(context *gin.Context) {
	cid := context.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "id must be int",
		})
		return
	}

	c := models.CalendarReminder{Id: id}
	err = c.Get()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"error":  err.Error(),
		})
		return
	}

	userId := getUserIdfromJWT(context)
	if userId != c.UserId {
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "no auth update",
		})
		return
	}

	c = models.CalendarReminder{Id: id}
	ra, err := c.Delete()

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "delete fail",
		})
		return
	}
	if ra == 0 {
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "person not exists",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "delete success",
		})
	}

}

func UpdateCalendarReminder(context *gin.Context) {
	cid := context.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "id must be int",
		})
		return
	}

	var param AddParam
	if err := context.ShouldBind(&param); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	layout := "2006-01-02 15:04:05" // Go的time包规定的日期格式字符串
	t, err := time.Parse(layout, param.ReminderTime)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "reminderTime参数校验错误，请输入正确的日期格式（YYYY-MM-dd HH:mm:ss）",
		})
		return
	}

	c := models.CalendarReminder{Id: id}
	err = c.Get()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"error":  err.Error(),
		})
		return
	}

	userId := getUserIdfromJWT(context)
	if userId != c.UserId {
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "no auth update",
		})
		return
	}

	c = models.CalendarReminder{Id: id, Content: param.ReminderContent, Time: t}
	ra, err := c.Update()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "update fail",
		})
		return
	}
	if ra == 0 {
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "not found",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "update success",
		})
	}

}
func ListCalendarReminder(context *gin.Context) {

	userId := getUserIdfromJWT(context)
	c := models.CalendarReminder{UserId: userId}

	rs, err := c.List()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "get list fail",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  0,
		"persons": rs,
	})
}

func getUserIdfromJWT(context *gin.Context) (userid int) {
	authHeader := context.Request.Header.Get("Authorization")
	if authHeader == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "未找到Authorization字段"})
		return
	}
	// 解析JWT令牌
	tokenString := authHeader[len("Bearer "):]
	token, err := utils.ParseToken(tokenString)
	if err != nil || !token.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"error": "令牌不合法"})
		return
	}

	// 获取负载信息
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": "令牌不合法"})
		return
	}

	userId, ok := claims["user"].(string)
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": "令牌中的用户不存在"})
		return
	}

	userid, err = strconv.Atoi(userId)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "userid error",
		})
	}
	return
}
