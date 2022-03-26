package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/database"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/domain/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/domain/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/routers"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/pkg/csv_helper"
)

func init() {
	database.Connect()

	authorRepository := author.NewAuthorRepository(database.DB)
	bookRepository := book.NewBookRepository(database.DB)

	author.InsertInitialAuthorData(authorRepository)
	book.InsertInitialBookData(bookRepository)

}

func main() {

	r := routers.SetupRouter()

	srv := &http.Server{
		Addr:         "0.0.0.0:8090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	log.Println("Server is running...")

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ShutdownServer(srv, time.Second*10)

}

func ShutdownServer(srv *http.Server, timeout time.Duration) {
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func InitiliazeDatabase() error {

	dbExist := database.DB.Migrator().HasTable(&book.Book{})

	//chech DB is already exist
	if !dbExist {
		database.DB.AutoMigrate(&book.Book{})
		database.DB.AutoMigrate(&author.Author{})

		// read CSV
		authorSlice, err := csv_helper.ReadCsv("../resources/author.csv", 1)

		if err != nil {
			fmt.Println(err)
			return errors.New("CSV cannot be read!")
		}

		//create author data
		var authorData []author.Author
		for _, author2 := range authorSlice {

			newAuthor := author.NewAuthor(author2[1])

			authorData = append(authorData, *newAuthor)

		}

		// add author data
		authorRepo := author.NewAuthorRepository(database.DB)
		authorRepo.InsertInitialData(authorData)

		//read book csv
		bookSlice, err := csv_helper.ReadCsv("../resources/book.csv", 1)

		if err != nil {
			fmt.Println(err)
			return errors.New("CSV cannot be read!")
		}

		//create book data
		var bookData []book.Book

		for _, book2 := range bookSlice {

			price, _ := strconv.ParseFloat(book2[4], 64)
			pageNum, _ := strconv.Atoi(book2[3])
			stockQuantity, _ := strconv.Atoi(book2[5])
			authorID, _ := strconv.Atoi(book2[6])
			newBook := book.NewBook(pageNum, stockQuantity, price, book2[0], book2[2], book2[1], uint(authorID))

			bookData = append(bookData, *newBook)

		}

		//insert book data
		bookRepo := book.NewBookRepository(database.DB)
		bookRepo.InsertInitialData(bookData)

	}

	return nil
}
