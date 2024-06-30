package main

import (
	"bytes"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	db "time-manager/utils"
)

func insertCalendarData() (id int64, err error) {
	//rs, err := db.SqlDB.Exec("delete from calendar_reminder where id =1")
	rs, err := db.SqlDB.Exec("insert into calendar_reminder (content,time,user_id) values(?,?,?)",
		"提醒测试123", "2021-06-30 12:00:00", 3)
	if err != nil {
		log.Println(err)
	}
	id, err = rs.LastInsertId()
	return
}

func TestAddCalendar(t *testing.T) {

	// 测试用例
	jsonValue1, _ := json.Marshal(map[string]string{"reminderContent": "日志提醒服务test", "reminderTime": "2024-06-30 12:00:00"})
	tests := []struct {
		name   string
		param  []byte
		expect string
	}{
		{"normal case", jsonValue1, "insert success"},
	}

	r := InitRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock一个HTTP请求
			req := httptest.NewRequest(
				"POST",                    // 请求方法
				"/calendar-reminder",      // 请求URL
				bytes.NewBuffer(tt.param), // 请求参数
			)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHB0aW1lIjoxNzE5NjY2MDIyLCJ1c2VyIjoiMyJ9.Hy19w9cm_sJtRlIOdnN78IQ898Gs1caMMSPJAXZJZEE")

			// mock一个响应记录器
			w := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)
			log.Println(w.Body.String())

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

func TestGet(t *testing.T) {

	// 测试用例
	//jsonValue1, _ := json.Marshal(map[string]string{"reminderContent": "日志提醒服务test", "reminderTime": "2024-06-30 12:00:00"})
	tests := []struct {
		name   string
		param  []byte
		expect string
	}{
		{"normal case", nil, "get success"},
	}

	r := InitRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock一个HTTP请求
			req := httptest.NewRequest(
				"GET",                     // 请求方法
				"/calendar-reminder",      // 请求URL
				bytes.NewBuffer(tt.param), // 请求参数
			)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHB0aW1lIjoxNzE5NjY2MDIyLCJ1c2VyIjoiMyJ9.Hy19w9cm_sJtRlIOdnN78IQ898Gs1caMMSPJAXZJZEE")

			// mock一个响应记录器
			w := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)
			log.Println(w.Body.String())

			// 校验状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)

			// 解析并检验响应内容是否复合预期
			var resp map[string]interface{}
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			assert.Nil(t, err)
		})
	}
}

func TestUpdateCalendar(t *testing.T) {
	id, _ := insertCalendarData()
	// 测试用例
	jsonValue1, _ := json.Marshal(map[string]string{"reminderContent": "日志提醒服务test", "reminderTime": "2024-06-30 12:00:00"})
	tests := []struct {
		name   string
		param  []byte
		expect string
	}{
		{"normal case", jsonValue1, "update success"},
	}

	r := InitRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock一个HTTP请求
			req := httptest.NewRequest(
				"PUT", // 请求方法
				"/calendar-reminder/"+strconv.FormatInt(id, 10), // 请求URL
				bytes.NewBuffer(tt.param),                       // 请求参数
			)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHB0aW1lIjoxNzE5NjY2MDIyLCJ1c2VyIjoiMyJ9.Hy19w9cm_sJtRlIOdnN78IQ898Gs1caMMSPJAXZJZEE")

			// mock一个响应记录器
			w := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)
			log.Println(w.Body.String())

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

func TestDelCalendar(t *testing.T) {
	id, _ := insertCalendarData()
	// 测试用例
	jsonValue1, _ := json.Marshal(map[string]string{"reminderContent": "日志提醒服务test", "reminderTime": "2024-06-30 12:00:00"})
	tests := []struct {
		name   string
		param  []byte
		expect string
	}{
		{"normal case", jsonValue1, "delete success"},
	}

	r := InitRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock一个HTTP请求
			req := httptest.NewRequest(
				"DELETE", // 请求方法
				"/calendar-reminder/"+strconv.FormatInt(id, 10), // 请求URL
				nil, // 请求参数
			)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHB0aW1lIjoxNzE5NjY2MDIyLCJ1c2VyIjoiMyJ9.Hy19w9cm_sJtRlIOdnN78IQ898Gs1caMMSPJAXZJZEE")

			// mock一个响应记录器
			w := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)
			log.Println(w.Body.String())

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
