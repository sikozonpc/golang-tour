package main

import "code.google.com/p/go-tour/pic"

// Pic : Creates a picture 
func Pic(dx, dy int) [][]uint8 {
	picSlice := make([][]uint8, dy)

	for y := range picSlice {

		picSlice[y] = make([]uint8, dx)

		for x := range picSlice[y] {

			picSlice[y][x] = uint8(x*10 ^ y*50)
		}

	}

	return picSlice
}

func main() {
	pic.Show(Pic)
}
