package book

import (
	"errors"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/domain/author"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID            uint          `gorm:"primarykey" json:"id"`
	Name          string        `json:"name"`
	StockCode     string        `json:"stock_code"`
	ISBN          string        `json:"isbn"`
	PageNum       int           `json:"page_num"`
	StockQuantity int           `json:"stock_quantity"`
	Price         float64       `json:"price"`
	AuthorID      uint          `json:"author_id"`
	Author        author.Author `gorm:"foreignKey:AuthorID" json:"author"`
}

//struct constructor
func NewBook(pageNum, stockNumber int, price float64, bookName, stockCode, isbn_num string, authorID uint) *Book {

	book := &Book{
		Name:          bookName,
		StockCode:     stockCode,
		ISBN:          isbn_num,
		PageNum:       pageNum,
		StockQuantity: stockNumber,
		Price:         price,
		AuthorID:      authorID,
	}

	return book
}

//This function check some rules and decrease the stock quantity
func (b *Book) Buy(count int) error {

	//check ıs there enough book to buy
	if count > b.StockQuantity {
		err := errors.New("Yeterli sayida kitap yoktur lutfen daha az miktarda deneyiniz!")
		return err
	}

	b.StockQuantity -= count

	return nil
}
