package main

import (
	"code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		picture[x] = make([]uint8, dy)
		for y := 0; y < dy; y++ {
			picture[x][y] = uint8(x * y)
		}
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
