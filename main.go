package main

import (
	db "time-manager/utils"
)

func main() {
	defer db.SqlDB.Close()
	router := InitRouter()
	router.Run(":8000")
}
