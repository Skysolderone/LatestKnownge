package test

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/skip2/go-qrcode"
)

// easy qrcode
func TestQrcode(t *testing.T) {
	err := qrcode.WriteFile("https://www.google.com", qrcode.Medium, 256, "testQrCode.png")
	if err != nil {
		t.Log(err)
	}
	t.Log("genarate qrcode")
}

// custom qrcode
func TestCustomQrcode(t *testing.T) {
	qr, err := qrcode.New("url or string or picture_url", qrcode.Medium)
	errProcess(err)
	qr.BackgroundColor = color.RGBA{255, 255, 255, 255} //背景色
	qr.ForegroundColor = color.RGBA{0, 0, 0, 255}
	err = qr.WriteFile(256, "customqrcode.png")
	errProcess(err)
	t.Log("custom genarate qrcode")
}

func errProcess(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
