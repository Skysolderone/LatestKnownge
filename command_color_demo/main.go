package main

import (
	"os"

	"github.com/fatih/color"
)

func main() {
	color.Red("This is red text")
	color.Blue("This is blue text")
	color.Yellow("This is yellow text")

	// 创建一个红色文本的颜色对象
	red := color.New(color.FgRed).Add(color.Underline)
	red.Println("This is an underlined red text")

	// 创建带有背景色的绿色文本
	greenBg := color.New(color.FgGreen).Add(color.BgWhite)
	greenBg.Println("This is a green text with white background")

	// 创建带有多种样式的文本
	boldBlue := color.New(color.FgBlue, color.Bold)
	boldBlue.Println("This is a bold blue text")

	// 将红色应用于标准输出
	reds := color.New(color.FgRed)
	redWriter := reds.FprintFunc()

	// 重定向 os.Stdout 到红色输出
	redWriter(os.Stdout, "This is red text\n")

	// 输出到文件中也是可以的
	file, _ := os.Create("output.txt")
	defer file.Close()
	redWriter(file, "This red text will be written to a file")
}
