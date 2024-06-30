package handlers

import (
	"crypto/md5"
	_ "crypto/md5"
	"database/sql"
	"encoding/hex"
	_ "encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time-manager/models"
	. "time-manager/utils"
)

type LoginParam struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func AddUser(context *gin.Context) {

	var param LoginParam
	if err := context.ShouldBind(&param); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash := md5.Sum([]byte(param.Password))
	encryptedData := hex.EncodeToString(hash[:])
	u := models.User{Username: param.Username, Password: encryptedData}

	exist, err := u.HasUser()

	if exist {
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "username already exists",
		})
		return
	}

	_, err = u.Add()

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "add user fail",
		})
		return
	}
	msg := fmt.Sprintf("register user success: %s", param.Username)
	context.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    msg,
	})
}

func Login(context *gin.Context) {

	var param LoginParam
	if err := context.ShouldBind(&param); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash := md5.Sum([]byte(param.Password))
	encryptedData := hex.EncodeToString(hash[:])
	u := models.User{Username: param.Username, Password: encryptedData}

	err := u.GetByUser()
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{
				"status": -1,
				"msg":    "username or password is wrong",
			})
			return
		}
		log.Println(err)
	}
	tokenString, err := CreateToken(strconv.Itoa(u.Id))
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "username or password is wrong",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": 0,
		"token":  tokenString,
	})
}
