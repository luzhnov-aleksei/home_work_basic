package main

import (
	"fmt"
	"os"
)

func main() {
	width, height := 0, 0
	fmt.Print("Введите ширину поля: ")
	_, err := fmt.Scanln(&width)
	if err != nil {
		fmt.Println("Ошибка ввода ширины поля:", err)
		os.Exit(1)
	}
	fmt.Print("Введите высоту поля: ")
	_, err = fmt.Scanln(&height)
	if err != nil {
		fmt.Println("Ошибка ввода высоты поля:", err)
		os.Exit(1)
	}
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if j%2 == 0 {
				fmt.Print(" #")
			} else {
				fmt.Print("# ")
			}
		}
		fmt.Println()
	}
}
