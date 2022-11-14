package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	yPow := make([][]uint8, dy)
	for y := range yPow {
		xPow := make([]uint8, dx)
		for x := range xPow {
			// xPow[x] = uint8(x)
			// xPow[x] = uint8(x + y)
			// xPow[x] = uint8((x + y) / 2)
			// xPow[x] = uint8(x * y)
			xPow[x] = uint8(x ^ y)
		}
		yPow[y] = xPow
	}
	return yPow
}

func main() {
	pic.Show(Pic)
}
