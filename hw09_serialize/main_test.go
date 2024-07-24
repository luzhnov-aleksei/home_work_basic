package main

import (
	"testing"

	"github.com/luzhnov-aleksei/hw09_serialize/protofile"
	"google.golang.org/protobuf/proto"
)

func TestSerializeBooks(t *testing.T) {
	books := []*protofile.Book{
		{Id: 1, Title: "Go Programming", Author: "John Doe", Year: 2023, Size: 300, Rate: 4.5},
		{Id: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2022, Size: 250, Rate: 4.0},
	}

	data, err := serializeBooks(books)
	if err != nil {
		t.Fatalf("Failed to serialize books: %v", err)
	}

	if len(data) == 0 {
		t.Error("Serialized data is empty")
	}
}

func TestDeserializeBooks(t *testing.T) {
	originalBooks := []*protofile.Book{
		{Id: 1, Title: "Go Programming", Author: "John Doe", Year: 2022, Size: 300, Rate: 4.5},
		{Id: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2022, Size: 250, Rate: 4.0},
	}

	data, err := serializeBooks(originalBooks)
	if err != nil {
		t.Fatalf("Failed to serialize books: %v", err)
	}

	deserializedBooks, err := deserializeBooks(data)
	if err != nil {
		t.Fatalf("Failed to deserialize books: %v", err)
	}

	if len(deserializedBooks) != len(originalBooks) {
		t.Errorf("Expected %d books, got %d", len(originalBooks), len(deserializedBooks))
	}

	for i, book := range deserializedBooks {
		if !proto.Equal(book, originalBooks[i]) {
			t.Errorf("Book %d does not match. Expected %+v, got %+v", i, originalBooks[i], book)
		}
	}
}

func TestDeserializeBooksInvalidData(t *testing.T) {
	data := []byte{0x00, 0x01, 0x02, 0x03}

	_, err := deserializeBooks(data)
	if err == nil {
		t.Error("Expected error during deserialization, got nil")
	}
}
