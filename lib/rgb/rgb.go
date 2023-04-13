package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

var (
	src = []uint32{84, 172, 94} // 微信绿
	dst = []uint32{222, 25, 25} // 印章红
)

func main() {
	read := gocv.IMRead("1.jpg", gocv.IMReadFlag(1))
	img, err := read.ToImage()
	if err != nil {
		panic(err)
	}
	var count = 0
	for i := 0; i < read.Cols(); i++ {
		for j := 0; j < read.Rows(); j++ {
			px := img.At(i, j)
			r, g, b, _ := px.RGBA()
			if r == src[0] && g == src[1] && b == src[2] {
				count++
			}
		}
	}
	fmt.Println(count)
}
