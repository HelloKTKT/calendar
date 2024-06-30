package main

import (
	"container/list"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	utils "time-manager/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域请求
	},
}

var l = list.New() //用于wsl监听数据

func main() {
	// 定时器每秒扫描数据库，到达执行时间进行提醒，可将数据库内容放入redis等缓存进行性能提升，这里简单demo不再复杂系统
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("执行定时任务:", time.Now())
				list, err := ListCalendarReminders()
				if err != nil {
					log.Println(err)
				}
				if list != nil {
					for _, c := range list {
						fmt.Println("到达提醒时间:", c.Time)
						fmt.Println("提醒内容:", c.Content)
						fmt.Println(c)
						//监听数据，推送到websocket /ws
						l.PushBack(c)
					}

					//进行邮件或者短信提醒,可在calendar_reminder加字段判断用哪种方式
					//utils.SendMsg()
					//utils.SendMail()
				}
			}
		}
	}()
	defer utils.SqlDB.Close()
	//gin
	router := InitRouter()
	router.Run(":8000")
}

// 每秒扫描数据库，到达执行时间进行提醒，可将数据库内容放入redis等缓存进行性能提升，这里简单demo不再复杂系统

type CalendarReminder struct {
	Id      int       `json:"id" form:"id"`
	Content string    `json:"content" form:"content"`
	Time    time.Time `json:"time" form:"time"`
	UserId  int       `json:"user_id" form:"user_id"`
}

func ListCalendarReminders() (calendarReminders []CalendarReminder, err error) {
	calendarReminders = make([]CalendarReminder, 0)
	rs, err := utils.SqlDB.Query("select * from calendar_reminder where time = now()")
	defer rs.Close()
	if err != nil {
		log.Println(err)
	}
	for rs.Next() {
		var cr CalendarReminder
		rs.Scan(&cr.Id, &cr.Content, &cr.Time, &cr.UserId)
		calendarReminders = append(calendarReminders, cr)
	}
	if err = rs.Err(); err != nil {
		return
	}
	return
}
