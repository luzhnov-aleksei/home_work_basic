package main

import (
	"fmt"
)

func main() {
	width,height := 0, 0
	fmt.Print("Введите ширину поля: ")
	fmt.Scanln(&width)
	fmt.Print("Введите высоту поля: ")
	fmt.Scanln(&height)
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
