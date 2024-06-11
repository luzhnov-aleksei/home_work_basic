package structcomparator

import (
	"errors"
)

type book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

// Set.
func (b *book) setID(id int) error {
	if id < 0 {
		return errors.New("некорректное id, введите положительное целое число")
	}
	b.id = id
	return nil
}

func (b *book) setTitle(title string) error {
	if len(title) > 99 {
		return errors.New("слишком длинное название книги, введите значение меньше 99 символов")
	}
	b.title = title
	return nil
}

func (b *book) setAuthor(author string) error {
	if len(author) > 99 {
		return errors.New("слишком длинное имя автора, введите значение меньше 99 символов")
	}
	b.author = author
	return nil
}

func (b *book) setYear(year int) error {
	if year <= 0 {
		return errors.New("некорректный год, введите положительное целое число")
	}
	b.year = year
	return nil
}

func (b *book) setSize(size int) error {
	if size < 0 {
		return errors.New("некорректный размер, введите положительное целое число")
	}
	b.size = size
	return nil
}

func (b *book) setRate(rate float64) error {
	if rate < 0 || rate > 5 {
		return errors.New("неккоректный рейтинг, введите значение от 0 до 5")
	}
	b.rate = rate
	return nil
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
