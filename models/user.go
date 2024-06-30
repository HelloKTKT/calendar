package models

import (
	"fmt"
	"log"
	db "time-manager/utils"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (u *User) Add() (id int64, err error) {
	rs, err := db.SqlDB.Exec("insert into user (username, password) values(?,?)",
		u.Username, u.Password)
	if err != nil {
		log.Println(err)
	}
	id, err = rs.LastInsertId()
	return
}

func (u *User) GetByUser() (err error) {
	err = db.SqlDB.QueryRow("select * from user where username=? and password=?", u.Username, u.Password).Scan(
		&u.Id, &u.Username, &u.Password)
	return
}

func (u *User) HasUser() (exist bool, err error) {
	rows, err := db.SqlDB.Query("select * from user where username=?", u.Username)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// 检查是否存在数据
	if rows.Next() {
		exist = true
	} else {
		exist = false
	}
	return
}
