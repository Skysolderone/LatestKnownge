package main

import (
	"image/png"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
)

//生成 CODE128

func main() {
	enc := oned.NewCode128Writer()
	img, _ := enc.Encode("Hello wws", gozxing.BarcodeFormat_CODE_128, 250, 50, nil)
	file, _ := os.Create("testCode.jpg")
	defer file.Close()
	_ = png.Encode(file, img)
}
