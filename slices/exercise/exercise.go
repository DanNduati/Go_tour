package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ss := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s[x] = uint8(x * y)
		}
		ss[y] = s
	}
	return ss
}

func main() {
	// Exercise : Implement Pic
	// It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
	pic.Show(Pic)
}
