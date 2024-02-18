package view

import (
	"bufio"
	"fmt"
	"github.com/cisnux-seed/book-shelf/entity"
	"github.com/cisnux-seed/book-shelf/repository"
	"github.com/cisnux-seed/book-shelf/utils"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var bookRepository = repository.BookRepository{}

func PrintMenu() {
	utils.ClearScreen()
	println("1. Add a new book")
	println("2. Show books")
	println("3. Update books")
	println("4. Delete books")
	print("What do you want to do? ")
	input, err := reader.ReadString('\n')
	if err != nil {
		println(err.Error())
		utils.Exit()
	}
	integerInput, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		println(err.Error())
		utils.Exit()
	}
	println()
	switch integerInput {
	case 1:
		utils.ClearScreen()
		err = addNewBookMenu()
		if err != nil {
			println(err.Error())
			utils.Exit()
		}
		isNavigatedUp, err := backToMainMenu()
		if err != nil {
			println(err.Error())
			utils.Exit()
		} else if isNavigatedUp {
			PrintMenu()
		} else {
			utils.Exit()
		}

	case 2:
		utils.ClearScreen()
		err = showBooks()
		if err != nil {
			println(err.Error())
			utils.Exit()
		}
		isNavigatedUp, err := backToMainMenu()
		if err != nil {
			println(err.Error())
			utils.Exit()
		} else if isNavigatedUp {
			PrintMenu()
		} else {
			utils.Exit()
		}
	case 3:
		utils.ClearScreen()
		err = updateBook()
		if err != nil {
			println(err.Error())
			utils.Exit()
		}
		isNavigatedUp, err := backToMainMenu()
		if err != nil {
			println(err.Error())
			utils.Exit()
		} else if isNavigatedUp {
			PrintMenu()
		} else {
			utils.Exit()
		}
	case 4:
		utils.ClearScreen()
		err = deleteBook()
		if err != nil {
			println(err.Error())
			utils.Exit()
		}
		isNavigatedUp, err := backToMainMenu()
		if err != nil {
			println(err.Error())
			utils.Exit()
		} else if isNavigatedUp {
			PrintMenu()
		} else {
			utils.Exit()
		}
	default:
		utils.Exit()
	}
}

func updateBook() (err error) {
	err = showBooks()
	if err != nil {
		return
	}
	print("What's your book's number? ")
	position, err := reader.ReadString('\n')
	integerPosition, err := strconv.Atoi(position[:len(position)-1])
	if err != nil {
		return
	}
	utils.ClearScreen()
	println("1. Update book by title")
	println("2. Update book by writer name")
	print("What do you want to do? ")
	selectedMenu, err := reader.ReadString('\n')
	integerSelectedMenu, err := strconv.Atoi(selectedMenu[:len(selectedMenu)-1])
	if err != nil {
		return
	}
	utils.ClearScreen()
	switch integerSelectedMenu {
	case 1:
		var title string
		print("Enter the new book's title! ")
		title, err = reader.ReadString('\n')
		title = title[:len(title)-1]
		if err != nil {
			return
		}
		err = bookRepository.UpdateAtTitle(integerPosition-1, title)
		if err != nil {
			return
		}
		utils.ClearScreen()
		err = showBooks()
		if err != nil {
			return
		}
	case 2:
		var writerName string
		print("Enter the new book's writer name! ")
		writerName, err = reader.ReadString('\n')
		writerName = writerName[:len(writerName)-1]
		if err != nil {
			return
		}
		err = bookRepository.UpdateAtWriterName(integerPosition-1, writerName)
		if err != nil {
			return
		}
		utils.ClearScreen()
		err = showBooks()
		if err != nil {
			return
		}
	default:
		utils.Exit()
	}
	return
}

func deleteBook() (err error) {
	println("1. Delete book by title")
	println("2. Delete book by writer name")
	print("What do you want to do? ")
	input, err := reader.ReadString('\n')
	integerInput, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		return
	}
	utils.ClearScreen()
	err = showBooks()
	if err != nil {
		return
	}
	switch integerInput {
	case 1:
		var title string
		print("Enter the book's title! ")
		title, err = reader.ReadString('\n')
		title = title[:len(title)-1]
		if err != nil {
			return
		}
		err = bookRepository.DeleteBookByTitle(title)
		if err != nil {
			return
		}
		utils.ClearScreen()
		err = showBooks()
		if err != nil {
			return
		}
	case 2:
		var writerName string
		print("Enter the book's writer name! ")
		writerName, err = reader.ReadString('\n')
		writerName = writerName[:len(writerName)-1]
		if err != nil {
			return
		}
		err = bookRepository.DeleteBookByWriterName(writerName)
		if err != nil {
			return
		}
		utils.ClearScreen()
		err = showBooks()
		if err != nil {
			return
		}
	default:
		utils.Exit()
	}
	return
}

func showBooks() (err error) {
	err, books := bookRepository.GetBooks()
	if err != nil {
		return
	}
	println("No\tTitle\t\t\tWriter Name")
	for index, book := range books {
		no := index + 1
		if len(book.Title()) >= 16 {
			fmt.Printf("%d\t%s\t%s\n", no, book.Title(), book.WriterName())
		} else {
			fmt.Printf("%d\t%s\t\t%s\n", no, book.Title(), book.WriterName())
		}
	}
	return
}

func addNewBookMenu() (err error) {
	print("Enter your book's title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	title = title[:len(title)-1]
	print("Enter your book's writer name: ")
	writerName, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	writerName = writerName[:len(writerName)-1]
	bookBuilder := entity.BookBuilder()
	bookBuilder.SetTitle(title)
	bookBuilder.SetWriterName(writerName)
	book := bookBuilder.Build()
	err = bookRepository.InsertNewBook(book)
	if err != nil {
		return
	}
	return
}

func backToMainMenu() (isNavigatedUp bool, err error) {
	print("Back to main menu?(yes/no): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	input = input[:len(input)-1]
	if strings.ToLower(input) == "yes" {
		isNavigatedUp = true
	}
	return
}
