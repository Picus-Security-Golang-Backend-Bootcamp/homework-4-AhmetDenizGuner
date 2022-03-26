package book

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/database"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/domain/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/pkg/csv_helper"
)

func GetAllBooks(bookRepository BookRepository) ([]Book, error) {

	books, err := bookRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return books, nil

}

func SearchBooks(key string, bookRepository BookRepository) ([]Book, error) {
	books, err := bookRepository.FindAllByKey(key)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func DeleteBook(id int, bookRepository BookRepository) error {
	err := bookRepository.DeleteById(id)

	if err != nil {
		return err
	}

	return nil

}

func BuyBook(id, amount int, bookRepository BookRepository) error {

	book, err := bookRepository.FindById(id)

	if err != nil {
		return err
	}

	err = book.Buy(amount)

	if err != nil {
		return err
	}

	err = bookRepository.Update(book)

	if err != nil {
		return err
	}

	return nil
}

func IncreaseBookStock(id, amount int, bookRepository BookRepository) error {
	book, err := bookRepository.FindById(id)

	if err != nil {
		return err
	}

	book.StockQuantity += amount

	err = bookRepository.Update(book)

	if err != nil {
		return err
	}

	return nil
}

func AddNewBook(book Book, bookRepository BookRepository, authorRepository author.AuthorRepository) error {

	//check author is exist
	_, err := author.GetAuthorByID(int(book.AuthorID), &authorRepository)

	if err != nil {
		return err
	}

	//check book id or name is exist

	fmt.Println("-----------")
	fmt.Println(book)

	err = bookRepository.Create(book)

	if err != nil {
		return err
	}

	return nil
}

func InsertInitialBookData(bookRepository *BookRepository) error {

	dbExist := database.DB.Migrator().HasTable(&Book{})

	if !dbExist {
		database.DB.AutoMigrate(&Book{})

		//read book csv
		bookSlice, err := csv_helper.ReadCsv("../resources/book.csv", 1)

		if err != nil {
			return errors.New("CSV cannot be read!")
		}

		//create book data
		var bookData []Book

		for _, book := range bookSlice {

			price, _ := strconv.ParseFloat(book[4], 64)
			pageNum, _ := strconv.Atoi(book[3])
			stockQuantity, _ := strconv.Atoi(book[5])
			authorID, _ := strconv.Atoi(book[6])
			newBook := NewBook(pageNum, stockQuantity, price, book[0], book[2], book[1], uint(authorID))

			bookData = append(bookData, *newBook)

		}

		//insert book data

		bookRepository.InsertInitialData(bookData)
	}

	return nil
}
