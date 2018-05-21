package main

import (
	dayu "github.com/holdno/alidayu"
	"fmt"
)

func main() {
	userInput := &dayu.UserParams{
		//AccessKeyId:   "阿里云的AccessKeyId",
		AccessKeyId:   "LTAIWF2peUpfIZCY",
		//AppSecret:     "阿里云的AppSecret",
		AppSecret:     "lyUXM1FXtIDeejNKirTydRHs9b3Cys",
		PhoneNumbers:  "15136859017",
		SignName:      "桦树林",
		TemplateCode:  "SMS_109525111",
		// 模板变量赋值，一定是json格式，注意转义
		TemplateParam: "{\"code\": \"123456\"}",
	}
	ok, msg, err := dayu.SendMessage(userInput)
	if ok {
		fmt.Println("短信发送成功")
	} else {
		// 根据业务进行错误处理
		fmt.Println(msg, err)
	}
}