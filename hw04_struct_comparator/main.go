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

// Set.
func (b *book) setID(id int) {
	b.id = id
}

func (b *book) setTitle(title string) {
	b.title = title
}

func (b *book) setAuthor(author string) {
	b.author = author
}

func (b *book) setYear(year int) {
	b.year = year
}

func (b *book) setSize(size int) {
	b.size = size
}

func (b *book) setRate(rate float64) {
	b.rate = rate
}

// Get.
func (b *book) getID() int {
	return b.id
}

func (b *book) getTitle() string {
	return b.title
}

func (b *book) getAuthor() string {
	return b.author
}

func (b *book) getYear() int {
	return b.year
}

func (b *book) getSize() int {
	return b.size
}

func (b *book) getRate() float64 {
	return b.rate
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

func (c Comparator) Compare(b1, b2 *book) bool {
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
	b1 := &book{
		id:     1,
		title:  "Землянка",
		author: "Лужнов А.В.",
		year:   2022,
		size:   289,
		rate:   4.9,
	}
	fmt.Println("Первая книга: ", b1)
	b2 := &book{1, "ЗемлЯнка 2", "Лужнов А.В.", 2024, 189, 5.0}
	b2.setTitle("Землянка 2")
	b2.setID(2)
	b2.setAuthor("Лужнов")
	b2.setRate(4.8)
	b2.setYear(2021)
	b2.setSize(188)
	fmt.Println(b2.getAuthor(), b2.getRate(), b2.getSize(), b2.getYear(), b2.getID(), b2.getTitle())
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
