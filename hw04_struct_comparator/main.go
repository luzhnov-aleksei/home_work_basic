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

func (b *book) setID(id int) {
	b.id = id
}

func (b *book) setTitle(title string) {
	b.title = title
}

func (b *book) bookID() int {
	return b.id
}

func (b *book) bookTitle() string {
	return b.title
}

type CompareMode int

const (
	ByYear CompareMode = iota
	BySize
	ByRate
)

type Comparator struct {
	mode CompareMode
}

func (c Comparator) Compare(b1, b2 book) bool {
	switch c.mode {
	case ByYear:
		return b1.year > b2.year
	case BySize:
		return b1.size > b2.size
	case ByRate:
		return b1.rate > b2.rate
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
	b2 := book{1, "ЗемлЯнка 2", "Лужнов А.В.", 2024, 189, 5.0}
	b2.setTitle("Землянка 2")
	b2.setID(2)
	fmt.Println(b2.bookID(), b2.bookTitle())
	fmt.Println("Вторая книга:")
	fmt.Println(b2.getBaseInfo())
	fmt.Println("\n----Сравнение книг----")
	yearComparator := Comparator{mode: ByYear}
	sizeComparator := Comparator{mode: BySize}
	rateComparator := Comparator{mode: ByRate}
	fmt.Printf("Сравнение по годам: %d > %d, Итог: %t \n", b1.year, b2.year, yearComparator.Compare(b1, b2))
	fmt.Printf("Сравнение по размеру : %d > %d, Итог: %t \n", b1.size, b2.size, sizeComparator.Compare(b1, b2))
	fmt.Printf("Сравнение по рейтингу : %g > %g, Итог: %t \n", b1.rate, b2.rate, rateComparator.Compare(b1, b2))
}
