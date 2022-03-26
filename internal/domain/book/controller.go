package book

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/domain/author"
	jsonhelper "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/pkg/json_helper"
	responsemodel "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/pkg/response_model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type BookHandler struct {
	db               *gorm.DB
	bookRepository   BookRepository
	authorRepository author.AuthorRepository
}

func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{
		db:               db,
		bookRepository:   *NewBookRepository(db),
		authorRepository: *author.NewAuthorRepository(db),
	}
}

func (bh *BookHandler) BookListHandler(w http.ResponseWriter, r *http.Request) {
	books, err := GetAllBooks(bh.bookRepository)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	d := responsemodel.ApiResponse{
		Data: books,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp, _ := json.Marshal(d)
	w.Write(resp)

}

func (bh *BookHandler) BookSearchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	books, err := SearchBooks(vars["search_item"], bh.bookRepository)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	var d responsemodel.ApiResponse

	//check result list is empty
	if len(books) == 0 {
		d = responsemodel.ApiResponse{
			Data: "There is no book with your search item, please try again using by differet search item/s!",
		}
	} else {
		d = responsemodel.ApiResponse{
			Data: books,
		}
	}

	resp, _ := json.Marshal(d)
	w.Write(resp)

}

func (bh *BookHandler) BookDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	err := DeleteBook(id, bh.bookRepository)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d := responsemodel.ApiResponse{
		Data: "Book ID: " + vars["id"] + "deleted.",
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp, _ := json.Marshal(d)
	w.Write(resp)
}

func (bh *BookHandler) BookBuyHandler(w http.ResponseWriter, r *http.Request) {

	if !r.URL.Query().Has("id") || !r.URL.Query().Has("amount") {
		http.Error(w, errors.New("Query params are not valid.").Error(), http.StatusBadRequest)
		return
	}

	bookId := r.URL.Query().Get("id")
	orderAmount := r.URL.Query().Get("amount")

	bookID, err := strconv.Atoi(bookId)
	orderAmountInt, err1 := strconv.Atoi(orderAmount)

	if err != nil || err1 != nil || bookID < 0 || orderAmountInt < 0 {
		http.Error(w, errors.New("Query params must be positive integer").Error(), http.StatusBadRequest)
		return
	}

	err2 := BuyBook(bookID, orderAmountInt, bh.bookRepository)

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadGateway)
	}

	d := responsemodel.ApiResponse{
		Data: "Book ID: " + bookId + " was bought. Order amount: " + orderAmount,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp, _ := json.Marshal(d)
	w.Write(resp)

}

func (bh *BookHandler) BookUpdateStockHandler(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("id") || !r.URL.Query().Has("amount") {
		http.Error(w, errors.New("Query params are not valid.").Error(), http.StatusBadRequest)
		return
	}

	bookId := r.URL.Query().Get("id")
	increaseAmount := r.URL.Query().Get("amount")

	bookID, err := strconv.Atoi(bookId)
	increaseAmountInt, err1 := strconv.Atoi(increaseAmount)

	if err != nil || err1 != nil || bookID < 0 || increaseAmountInt < 0 {
		http.Error(w, errors.New("Query params must be positive integer").Error(), http.StatusBadRequest)
		return
	}

	err = IncreaseBookStock(bookID, increaseAmountInt, bh.bookRepository)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d := responsemodel.ApiResponse{
		Data: "Book ID: " + bookId + " was updated. Stock increase amount: " + increaseAmount,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp, _ := json.Marshal(d)
	w.Write(resp)
}

func (bh *BookHandler) BookAddHandler(w http.ResponseWriter, r *http.Request) {
	var book Book

	err := jsonhelper.DecodeJSONBody(w, r, &book)
	if err != nil {
		var mr *jsonhelper.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	err = AddNewBook(book, bh.bookRepository, bh.authorRepository)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	d := responsemodel.ApiResponse{
		Data: "Book ID: " + book.Name + " added",
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp, _ := json.Marshal(d)
	w.Write(resp)

}
