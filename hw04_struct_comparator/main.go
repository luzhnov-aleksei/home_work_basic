package main

import "fmt"

type book struct {
	id     int
	title  string // логичнее было бы расположить строки в конце, для выравнивания структуры
	author string // но задание требует именно такой порядок
	year   int
	size   int
	rate   float64
}

func (b *book) getBaseInfo() string {
	return "Название книги: " + b.title + "\nАвтор книги: " + b.author
}

func (b *book) setAll(id int, title, author string, year, size int, rate float64) {
	b.id = id
	b.title = title
	b.author = author
	b.year = year
	b.size = size
	b.rate = rate
}

type CompareMode int

const (
	ByYear CompareMode = iota
	BySize
	ByRate
)

func (b book) Compare(other book, mode CompareMode) bool {
	switch mode {
	case ByYear:
		return b.year > other.year
	case BySize:
		return b.size > other.size
	case ByRate:
		return b.rate > other.rate
	default:
		return false
	}
}

func main() {
	b1 := book{
		id:     1,
		title:  "Землянка",
		author: "Лужнов А.В.",
		year:   2022,
		size:   289,
		rate:   4.9,
	}
	fmt.Println("Первая книга: ", b1)
	b2 := book{}
	b2.setAll(2, "Землянка 2", "Лужнов А.В.", 2024, 189, 5.0)
	fmt.Println("Вторая книга:")
	fmt.Println(b2.getBaseInfo())
	fmt.Println("\n----Сравнение книг----")
	fmt.Printf("Сравнение по годам: %d > %d, Итог: %t \n", b1.year, b2.year, b1.Compare(b2, ByYear))
	fmt.Printf("Сравнение по размеру : %d > %d, Итог: %t \n", b1.size, b2.size, b1.Compare(b2, BySize))
	fmt.Printf("Сравнение по рейтингу : %g > %g, Итог: %t \n", b1.rate, b2.rate, b2.Compare(b2, ByRate))
}
