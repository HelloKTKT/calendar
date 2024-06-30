package utils

import (
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strconv"
	"strings"
)

func SendMail() {
	/*
		IMAP/SMTP 设置方法
		用户名/帐户： 你的QQ邮箱完整的地址
		密码： 生成的授权码
		电子邮件地址： 你的QQ邮箱的完整邮件地址
		接收邮件服务器： imap.qq.com，使用SSL，端口号993
		发送邮件服务器： smtp.qq.com，使用SSL，端口号465或587
	*/
	// 邮件发送者、接收者、SMTP服务器配置
	fromAcc := "xxxxx@qq.com" //发送者邮箱
	fromUser := "shi"         //发送者名称(随意)
	password := "xxxxxxxxxx"  // 这不是你的邮箱密码，而是开启SMTP服务后获得的授权码
	smtpServer := "smtp.qq.com"
	smtpPort := 587 // 或465，具体取决于SMTP服务器配置
	// 构建认证信息
	auth := smtp.PlainAuth("", fromAcc, password, smtpServer)

	// 收发件人信息
	from := mail.Address{fromUser, fromAcc}
	to := mail.Address{"收件人的名字(可以随便)", "xxxxx@qq.com"}

	// 构建邮件头和正文
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = "Go使用 net/smtp发送邮件" //主题
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=UTF-8"

	var headerBuffer strings.Builder
	for k, v := range headers {
		headerBuffer.WriteString(k + ": " + v + "\r\n")
	}
	headerBuffer.WriteString("\r\n") // 空行分隔邮件头和邮件体
	body := "这是邮件的正文内容。"
	message := headerBuffer.String() + body

	// 发送邮件
	err := smtp.SendMail(smtpServer+":"+strconv.Itoa(smtpPort),
		auth,
		fromAcc,
		[]string{to.Address},
		[]byte(message))
	if err != nil {
		log.Fatalf("发送邮件失败: %v", err)
	}
	fmt.Println("发送邮件成功!")
}
