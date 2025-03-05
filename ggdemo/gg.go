package main

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
)

func main() {
	// 表格数据
	data := [][]string{
		{"PRICE", "QUANTITY"},
		{"12314", "2321414215"},
		{"Bob", "30"},
		{"Charlie", "28"},
	}

	// 表格参数
	cellWidth := 100.0
	cellHeight := 40.0
	titleHeight := 60.0 // 标题区域高度
	rows := len(data)
	cols := len(data[0])
	width := cellWidth * float64(cols)
	height := cellHeight*float64(rows) + titleHeight // 画布高度增加标题区域

	// 创建画布
	dc := gg.NewContext(int(width), int(height))
	dc.SetColor(color.White)
	dc.Clear()

	// **加载本地图字体**
	fontPath := "/System/Library/Fonts/Supplemental/Arial.ttf" // macOS 示例
	// fontPath := "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf" // Linux 示例
	// fontPath := "C:/Windows/Fonts/arial.ttf" // Windows 示例

	// 设置标题字体
	err := dc.LoadFontFace(fontPath, 24) // 标题使用较大字号
	if err != nil {
		fmt.Println("字体加载失败:", err)
		return
	}

	// **绘制标题**
	dc.SetColor(color.Black)
	title := "买入"
	dc.DrawStringAnchored(title, width/2, titleHeight/2, 0.5, 0.5)

	// **绘制表格线**
	dc.SetColor(color.Black)
	dc.SetLineWidth(2)
	for r := 0; r <= rows; r++ {
		y := titleHeight + float64(r)*cellHeight
		dc.DrawLine(0, y, width, y)
		dc.Stroke()
	}
	for c := 0; c <= cols; c++ {
		x := float64(c) * cellWidth
		dc.DrawLine(x, titleHeight, x, height)
		dc.Stroke()
	}

	// **设置表格字体**
	err = dc.LoadFontFace(fontPath, 16) // 普通单元格使用较小字号
	if err != nil {
		fmt.Println("字体加载失败:", err)
		return
	}

	// **填充表格数据**
	for r, row := range data {
		for c, text := range row {
			x := float64(c)*cellWidth + cellWidth/2
			y := titleHeight + float64(r)*cellHeight + cellHeight/2
			dc.DrawStringAnchored(text, x, y, 0.5, 0.5)
		}
	}

	// 保存图片
	dc.SavePNG("买入.png")
}
