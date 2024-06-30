package main

import (
	"bytes"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestRegister(t *testing.T) {
	_, err := deleteData()
	// 模拟SQL数据库
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// 模拟SQL查询
	mock.ExpectQuery("select * from user where username=?").
		WithArgs("test123").
		WillReturnRows(sqlmock.NewRows(nil))
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
