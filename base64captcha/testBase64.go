package main

import (
	"fmt"

	"github.com/mojocn/base64Captcha"
)

func main() {
	configDigit := base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}
	// 生成数字验证码
	id, captcha := base64Captcha.GenerateCaptcha("", configDigit)
	// 将验证码的ID和Base64编码的图像数据返回给客户端
	fmt.Println(id, captcha)
}
