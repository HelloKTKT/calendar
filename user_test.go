package main

import (
	"bytes"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	db "time-manager/utils"
)

func deleteData() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("delete from user")
	if err != nil {
		log.Println(err)
	}
	ra, err = rs.RowsAffected()
	return
}

func insertData() (id int64, err error) {
	rs, err := db.SqlDB.Exec("insert into user (username, password) values(?,?)",
		"test", "e10adc3949ba59abbe56e057f20f883e")
	if err != nil {
		log.Println(err)
	}
	id, err = rs.LastInsertId()
	return
}

func TestRegister(t *testing.T) {
	deleteData()

	// 测试用例
	jsonValue1, _ := json.Marshal(map[string]string{"username": "test123", "password": "test"})
	tests := []struct {
		name   string
		param  []byte
		expect string
	}{
		{"normal case", jsonValue1, "register user success: test123"},
		{"bad case: User exists", jsonValue1, "username already exists"},
	}

	r := InitRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock一个HTTP请求
			req := httptest.NewRequest(
				"POST",                    // 请求方法
				"/register",               // 请求URL
				bytes.NewBuffer(tt.param), // 请求参数
			)
			req.Header.Set("Content-Type", "application/json")

			// mock一个响应记录器
			w := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)

			// 校验状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)

			// 解析并检验响应内容是否复合预期
			var resp map[string]interface{}
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp["msg"])
		})
	}
}

func TestLogin(t *testing.T) {
	insertData()

	// 测试用例
	jsonValue1, _ := json.Marshal(map[string]string{"username": "test", "password": "123456"})
	tests := []struct {
		name   string
		param  []byte
		expect string
	}{
		{"normal case", jsonValue1, "register user success: test123"},
		{"bad case: User exists", jsonValue1, "username already exists"},
	}

	r := InitRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock一个HTTP请求
			req := httptest.NewRequest(
				"POST",                    // 请求方法
				"/login",                  // 请求URL
				bytes.NewBuffer(tt.param), // 请求参数
			)
			req.Header.Set("Content-Type", "application/json")

			// mock一个响应记录器
			w := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)

			// 校验状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)

			// 解析并检验响应内容是否复合预期
			var resp map[string]interface{}
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			assert.Nil(t, err)
			assert.NotEmpty(t, resp["token"])
		})
	}
}
