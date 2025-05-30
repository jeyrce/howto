package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	// 创建画布 (400x400)
	img := image.NewRGBA(image.Rect(0, 0, 400, 400))
	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.RGBA{R: 250, G: 240, B: 230, A: 255}}, image.Point{}, draw.Src) // 米色背景

	// 绘制粽体 (绿色三角形)
	drawTriangle(img, color.RGBA{R: 80, G: 120, B: 60, A: 255}, image.Point{X: 200, Y: 100}, 150, 200)

	// 绘制粽叶纹理 (深绿色线条)
	drawLeafVeins(img, color.RGBA{R: 50, G: 90, B: 40, A: 255}, image.Point{X: 200, Y: 100}, 150, 200)

	// 绘制捆绳 (红色交叉线)
	drawRope(img, color.RGBA{R: 180, G: 60, B: 50, A: 255}, image.Point{X: 200, Y: 150}, 100)

	// 保存为PNG
	f, _ := os.Create("zongzi.png")
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}

// 绘制三角形粽体 [2,6](@ref)
func drawTriangle(img *image.RGBA, c color.Color, center image.Point, width, height int) {
	points := []image.Point{
		{X: center.X, Y: center.Y - height/2},           // 顶点
		{X: center.X - width/2, Y: center.Y + height/2}, // 左下
		{X: center.X + width/2, Y: center.Y + height/2}, // 右下
	}
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			if inTriangle(x, y, points) {
				img.Set(x, y, c)
			}
		}
	}
}

// 绘制粽叶纹理 [3,4](@ref)
func drawLeafVeins(img *image.RGBA, c color.Color, center image.Point, width, height int) {
	for i := 0; i < 5; i++ {
		yOffset := center.Y - height/3 + i*20
		drawLine(img, center.X-width/3, yOffset, center.X+width/3, yOffset, c)
	}
}

// 绘制捆绳 [3](@ref)
func drawRope(img *image.RGBA, c color.Color, center image.Point, length int) {
	// 横向捆绳
	drawLine(img, center.X-length, center.Y, center.X+length, center.Y, c)
	// 纵向捆绳
	drawLine(img, center.X, center.Y-length/2, center.X, center.Y+length/2, c)
}

// --- 以下为工具函数 ---
// 判断点是否在三角形内 (扫描线算法)
func inTriangle(x, y int, points []image.Point) bool {
	v0 := points[0]
	v1 := points[1]
	v2 := points[2]
	d1 := (x-v1.X)*(v0.Y-v1.Y) - (y-v1.Y)*(v0.X-v1.X)
	d2 := (x-v2.X)*(v1.Y-v2.Y) - (y-v2.Y)*(v1.X-v2.X)
	d3 := (x-v0.X)*(v2.Y-v0.Y) - (y-v0.Y)*(v2.X-v0.X)
	return (d1 >= 0 && d2 >= 0 && d3 >= 0) || (d1 <= 0 && d2 <= 0 && d3 <= 0)
}

// 绘制直线 (Bresenham算法) [2](@ref)
func drawLine(img *image.RGBA, x0, y0, x1, y1 int, c color.Color) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 > x1 {
		sx = -1
	}
	if y0 > y1 {
		sy = -1
	}
	err := dx - dy
	for {
		img.Set(x0, y0, c)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
