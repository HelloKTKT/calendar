package models

import (
	"log"
	"time"
	db "time-manager/utils"
)

type CalendarReminder struct {
	Id      int       `json:"id" form:"id"`
	Content string    `json:"content" form:"content"`
	Time    time.Time `json:"time" form:"time"`
	UserId  int       `json:"user_id" form:"user_id"`
}

func (c *CalendarReminder) Add() (id int64, err error) {
	rs, err := db.SqlDB.Exec("insert into calendar_reminder (content, time, user_id) values(?,?,?)",
		c.Content, c.Time, c.UserId)
	if err != nil {
		log.Println(err)
	}
	id, err = rs.LastInsertId()
	return
}

func (c *CalendarReminder) Delete() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("delete from calendar_reminder where id=?", c.Id)
	if err != nil {
		log.Println(err)
	}
	ra, err = rs.RowsAffected()
	return
}

func (c *CalendarReminder) Update() (ra int64, err error) {
	stmt, err := db.SqlDB.Prepare("update calendar_reminder set content=?,time=? where id=?")
	if err != nil {
		log.Println(err)
	}
	rs, err := stmt.Exec(c.Content, c.Time, c.Id)
	if err != nil {
		log.Println(err)
	}
	ra, err = rs.RowsAffected()
	return
}

func (c *CalendarReminder) Get() (err error) {
	err = db.SqlDB.QueryRow("select * from calendar_reminder where id=?", c.Id).Scan(
		&c.Id, &c.Content, &c.Time, &c.UserId)
	return
}

func (c *CalendarReminder) List() (calendarReminders []CalendarReminder, err error) {
	calendarReminders = make([]CalendarReminder, 0)
	rs, err := db.SqlDB.Query("select * from calendar_reminder where user_id=?", c.UserId)
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
