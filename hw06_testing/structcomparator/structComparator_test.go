package structcomparator

import (
	"testing"
)

func TestSetters(t *testing.T) {
	tests := []struct {
		name     string
		setFunc  func(b *book) error
		expected interface{}
		wantErr  bool
	}{
		{"Установка названия", func(b *book) error {
			return b.setTitle("Хорошая книга")
		}, "Хорошая книга", false},
		{"Установка названия", func(b *book) error {
			return b.setTitle(`Очень Очень Очень Очень Очень Очень Очень Очень Очень
			Очень Очень Очень Очень Очень Очень Длинное Название`)
		}, "", true},
		{"Установка автора", func(b *book) error { return b.setAuthor("Иван Иванов") }, "Иван Иванов", false},
		{"Установка автора", func(b *book) error {
			return b.setAuthor(`Очень Очень Очень Очень Очень Очень Очень Очень Очень
			Очень Очень Очень Очень Очень Очень Длинное Имя Автора`)
		}, "", true},
		{"Установка года", func(b *book) error { return b.setYear(2020) }, 2020, false},
		{"Установка года", func(b *book) error { return b.setYear(-1) }, 0, true},
		{"Установка размера", func(b *book) error { return b.setSize(100) }, 100, false},
		{"Установка размера", func(b *book) error { return b.setSize(-1) }, 0, true},
		{"Установка рейтинга", func(b *book) error { return b.setRate(4.5) }, 4.5, false},
		{"Установка рейтинга", func(b *book) error { return b.setRate(6) }, 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &book{}
			err := tt.setFunc(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("ожидается ошибка: %v, получено: %v", tt.wantErr, err)
			}
			v := tt.expected
			if tt.name == "Установка корректного года" && b.year != v {
				t.Errorf("ожидается год: %d, получено: %d", v, b.year)
			}
			if tt.name == "Установка корректного размера" && b.size != v {
				t.Errorf("ожидается размер: %d, получено: %d", v, b.size)
			}
			if tt.name == "Установка корректного названия" && b.title != v {
				t.Errorf("ожидается название: %s, получено: %s", v, b.title)
			}
			if tt.name == "Установка корректного автора" && b.author != v {
				t.Errorf("ожидается автор: %s, получено: %s", v, b.author)
			}
			if tt.name == "Установка корректного рейтинга" && b.rate != v {
				t.Errorf("ожидается рейтинг: %f, получено: %f", v, b.rate)
			}
		})
	}
}

func TestGetters(t *testing.T) {
	b := &book{
		title:  "Тестовая книга",
		author: "Иван Иванов",
		year:   2021,
		size:   123,
		rate:   4.5,
	}
	tests := []struct {
		name     string
		getFunc  func() interface{}
		expected interface{}
	}{
		{"Получение названия", func() interface{} { return b.getTitle() }, "Тестовая книга"},
		{"Получение автора", func() interface{} { return b.getAuthor() }, "Иван Иванов"},
		{"Получение года", func() interface{} { return b.getYear() }, 2021},
		{"Получение размера", func() interface{} { return b.getSize() }, 123},
		{"Получение рейтинга", func() interface{} { return b.getRate() }, 4.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.getFunc(); got != tt.expected {
				t.Errorf("ожидается: %v, получено: %v", tt.expected, got)
			}
		})
	}
}

func TestComparator(t *testing.T) {
	b1 := &book{year: 2020, size: 300, rate: 4.5}
	b2 := &book{year: 2021, size: 250, rate: 4.0}

	tests := []struct {
		name     string
		mode     CompareMode
		expected bool
	}{
		{"Сравнение по году", ByYear, false},
		{"Сравнение по размеру", BySize, true},
		{"Сравнение по рейтингу", ByRate, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comparator := Comparator{mode: tt.mode}
			result := comparator.Compare(b1, b2)
			if result != tt.expected {
				t.Errorf("ожидается: %v, получено: %v", tt.expected, result)
			}
		})
	}
}
