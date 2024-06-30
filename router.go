package main

import (
	"container/list"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
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

	// WebSocket服务
	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer conn.Close()
		//监听list
		for {
			// 创建一个每秒触发一次的计时器
			ticker := time.NewTicker(time.Second)
			preLen := l.Len()
			// 无限循环，以便每次计时器触发时执行代码，list有增量则有数据更新 并进行推送
			for range ticker.C {
				fmt.Println("data change check")
				if l.Len() > 0 && l.Len() > preLen { //数据变化说明有新的推送,才进行推送
					fmt.Println("data change")
					nowLen := l.Len()
					//nowList := l
					//array := listToSlice(nowList)
					for i := preLen; i < nowLen; i++ {
						fmt.Println("data change123123")
						conn.WriteMessage(websocket.TextMessage, []byte("新的时间提醒"))
					}
					preLen = nowLen
				}
			}

		}
	})
	return router
}

func listToSlice(l *list.List) []interface{} {
	slice := make([]interface{}, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		slice = append(slice, e.Value)
	}
	return slice
}
