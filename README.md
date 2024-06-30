## 项目
基于Gin的Go语言简单学习例子，日历提醒系统
<br>
需要注册账号，登陆后可使用日期提醒管理
<br>
单元测试只加了几个简单的基本的normal case
<br>
覆盖率：coverage: 72.7% of statements
ok      time-manager    0.538s

## go安装
### Mac上安装
```shell
brew install go
```

### 环境变量配置
在~/.bashrc 或则 ~/.profile中添加
```shell
export GOROOT=/usr/local/Cellar/go/1.10/libexec  //具体参考自己安装的版本
export GOPATH=/Users/lfuture/Documents/goProject/go-rest:/Users/lfuture/Documents/goProject/其他项目地址 // 添加项目地址,第一个目录
填写你即将clone项目在你本地保存的目录
export GOBIN=/Users/lfuture/bin
```
使配置生效
```shell
source ~/.bashrc // 或则 source ~/.profile
```

## 配置及运行
```shell
git clone git@github.com:lfuture/go-rest.git
cd go-rest
go get github.com/gin-gonic/gin
```
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
修改`mydatabase/mysql.go`
```go
SqlDB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/calendarSys?parseTime=true")
```
### 运行
```shell
go run main.go router.go //直接运行
go build                // 打包
```
### 访问
即可访问`router.go`里定义的路由#