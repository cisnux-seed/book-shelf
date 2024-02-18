package repository

import (
	"bufio"
	"cmp"
	"errors"
	"fmt"
	"github.com/cisnux-seed/book-shelf/entity"
	"io"
	"os"
	"slices"
	"strings"
)

var (
	NotFoundError = errors.New("book is not found")
)

const fileName = "book_shelf.txt"

func init() {
	file, err := os.OpenFile(fileName, os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
}

type BookRepository struct {
	books []entity.Book
}

func (bookRepository *BookRepository) InsertNewBook(book *entity.Book) (err error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()
	stringBook := fmt.Sprintf("%s,%s,%s\n", book.Id(), book.Title(), book.WriterName())
	_, err = file.WriteString(stringBook)
	return
}

func (bookRepository *BookRepository) InsertNewBooks(books []entity.Book) (err error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()
	for _, book := range books {
		stringBook := fmt.Sprintf("%s,%s,%s\n", book.Id(), book.Title(), book.WriterName())
		_, err = file.WriteString(stringBook)
	}
	return
}

func (bookRepository *BookRepository) DeleteBookByTitle(title string) (err error) {
	slices.SortFunc(bookRepository.books, func(a, b entity.Book) int {
		return cmp.Compare(a.Title(), b.Title())
	})
	position, isFound := slices.BinarySearchFunc(bookRepository.books, title, func(book entity.Book, title string) int {
		return cmp.Compare(book.Title(), title)
	})
	if isFound {
		newBooks := make([]entity.Book, len(bookRepository.books)-1, cap(bookRepository.books)-1)
		for i, j := 0, 0; i < len(bookRepository.books); i++ {
			if position == i {
				continue
			}
			newBooks[j] = bookRepository.books[i]
			j++
		}
		if len(newBooks) > 0 {
			err = bookRepository.InsertNewBooks(newBooks)
			if err != nil {
				return
			}
		}
	}
	return
}

func (bookRepository *BookRepository) DeleteBookByWriterName(writerName string) (err error) {
	slices.SortFunc(bookRepository.books, func(a, b entity.Book) int {
		return cmp.Compare(a.WriterName(), b.WriterName())
	})
	position, isFound := slices.BinarySearchFunc(bookRepository.books, writerName, func(book entity.Book, writerName string) int {
		return cmp.Compare(book.WriterName(), writerName)
	})
	if isFound {
		newBooks := make([]entity.Book, len(bookRepository.books)-1, cap(bookRepository.books)-1)
		for i, j := 0, 0; i < len(bookRepository.books); i++ {
			if position == i {
				continue
			}
			newBooks[j] = bookRepository.books[i]
			j++
		}
		if len(newBooks) > 0 {
			err = bookRepository.InsertNewBooks(newBooks)
			if err != nil {
				return
			}
		}
	}
	return
}

func (bookRepository *BookRepository) UpdateAtTitle(position int, newTitle string) (err error) {
	if position >= len(bookRepository.books) {
		err = NotFoundError
		return
	}

	bookBuilder := entity.BookBuilder()
	bookBuilder.SetId(bookRepository.books[position].Id())
	bookBuilder.SetTitle(newTitle)
	bookBuilder.SetWriterName(bookRepository.books[position].WriterName())
	bookRepository.books[position] = *bookBuilder.Build()

	err = bookRepository.InsertNewBooks(bookRepository.books)
	if err != nil {
		return
	}
	return
}

func (bookRepository *BookRepository) UpdateAtWriterName(position int, newWriterName string) (err error) {
	if position >= len(bookRepository.books) {
		err = NotFoundError
		return
	}

	bookBuilder := entity.BookBuilder()
	bookBuilder.SetId(bookRepository.books[position].Id())
	bookBuilder.SetTitle(bookRepository.books[position].Title())
	bookBuilder.SetWriterName(newWriterName)
	bookRepository.books[position] = *bookBuilder.Build()

	err = bookRepository.InsertNewBooks(bookRepository.books)
	if err != nil {
		return
	}
	return
}

func (bookRepository *BookRepository) GetBooks() (err error, books []entity.Book) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()
	reader := bufio.NewReader(file)
	for line, _, err := reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
		stringBook := strings.Split(string(line), ",")
		bookBuilder := entity.BookBuilder()
		bookBuilder.SetId(stringBook[0])
		bookBuilder.SetTitle(stringBook[1])
		bookBuilder.SetWriterName(stringBook[2])
		book := bookBuilder.Build()
		books = append(books, *book)
	}
	bookRepository.books = books
	return
}
