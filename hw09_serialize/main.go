package main

import (
	"fmt"
	"log"

	"github.com/luzhnov-aleksei/hw09_serialize/protofile"
	"google.golang.org/protobuf/proto"
)

func serializeBooks(books []*protofile.Book) ([]byte, error) {
	bookList := &protofile.BookList{Books: books}
	return proto.Marshal(bookList)
}

func deserializeBooks(data []byte) ([]*protofile.Book, error) {
	var bookList protofile.BookList
	err := proto.Unmarshal(data, &bookList)
	if err != nil {
		return nil, err
	}
	return bookList.Books, nil
}

func main() {
	books := []*protofile.Book{
		{Id: 1, Title: "Go Programming", Author: "John Doe", Year: 2023, Size: 300, Rate: 4.5},
		{Id: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2022, Size: 250, Rate: 4.0},
	}

	data, err := serializeBooks(books)
	if err != nil {
		log.Fatalf("Failed to serialize books: %v", err)
	}

	deserializedBooks, err := deserializeBooks(data)
	if err != nil {
		log.Fatalf("Failed to deserialize books: %v", err)
	}

	for _, book := range deserializedBooks {
		fmt.Printf("%+v\n", book)
	}
}
