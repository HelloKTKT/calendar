package utils

import (
	gosms "github.com/pkg6/go-sms"
	"github.com/pkg6/go-sms/gateways"
	"github.com/pkg6/go-sms/gateways/twilio"
)

func SendMsg() {
	sms := gosms.NewParser(gateways.Gateways{Twilio: twilio.Twilio{AccountSID: "ACd********", AuthToken: "***********", TwilioPhoneNumber: "+1********"}})
	// 常规
	sms.Send(18888888888, gosms.MapStringAny{
		"content":  "你的日志提醒服务：****。",
		"template": "SMS_001",
		"data": gosms.MapStrings{
			"code": "6379",
		},
	}, nil)
}
