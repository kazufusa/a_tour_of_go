package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) (pic [][]uint8) {
    for x := 0; x < dx; x++ {
        pic = append(pic, make([]uint8, dy))
        for y := 0; y < dy; y++ {
            pic[x][y] = uint8(x ^ y)
        }
    }
    return
}

func main() {
    pic.Show(Pic)
}
