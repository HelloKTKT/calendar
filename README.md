## 项目
基于Gin的Go语言简单学习例子，日历提醒系统
<br>
需要注册账号,登陆后可使用日期提醒管理,只能管理自己的提醒信息
<br>
单元测试只加了几个简单的基本的normal case
<br>
覆盖率：coverage: 72.7% of statements
ok      time-manager    0.538s
<br>
到达日期提醒时间发短信或者邮件进行提醒（只进行简单封装，需要密钥才能实现）
websocket（已实现）


### 创建数据库
```mysql
create database calendarSys;
CREATE TABLE `user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(60) NOT NULL DEFAULT '',
    `password` varchar(60) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET=utf8;

CREATE TABLE `calendar_reminder` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` varchar(600) NOT NULL DEFAULT '',
  `time` DATETIME NOT NULL,
  `user_id` int(11),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET=utf8;
```
### 修改数据库
修改`utils/mysql.go`
```go
SqlDB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/calendarSys?parseTime=true")
```
### 运行
```shell
go run main.go router.go //直接运行
```
### 访问
即可访问`router.go`里定义的路由#