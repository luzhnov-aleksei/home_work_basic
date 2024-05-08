package main

import (
	"fmt"
)

func main() {
	width := 5
	height := 5
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if j%2 == 0 {
				fmt.Print(" #")
			} else {
				fmt.Print("# ")
			}
		}
		fmt.Print("\n")
	}

}
