package routers

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/database"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/domain/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/domain/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	r := mux.NewRouter()

	CORSOptions()

	bookHandler := book.NewBookHandler(database.DB)
	authorHandler := author.NewAuthorHandler(database.DB)

	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.AuthenticationMiddleware)

	bookRoute := r.PathPrefix("/book").Subrouter()
	authotRoute := r.PathPrefix("/author").Subrouter()

	bookRoute.HandleFunc("/list", bookHandler.BookListHandler).Methods("GET")
	bookRoute.HandleFunc("/search/{search_item}", bookHandler.BookSearchHandler).Methods("GET")
	bookRoute.HandleFunc("/buy", bookHandler.BookBuyHandler).Methods("GET")
	bookRoute.HandleFunc("/delete/{id:[0-9]+}", bookHandler.BookDeleteHandler).Methods("GET")
	bookRoute.HandleFunc("/update", bookHandler.BookUpdateStockHandler).Methods("GET")
	bookRoute.HandleFunc("/add", bookHandler.BookAddHandler).Methods("POST")

	authotRoute.HandleFunc("/get{id:[0-9]+}", authorHandler.GetAuthorByIdHandler).Methods("GET")
	authotRoute.HandleFunc("/search/{name}", authorHandler.GetAuthorByNameHandler).Methods("GET")
	authotRoute.HandleFunc("/add", authorHandler.AddAuthorHandler).Methods("GET")

	return r

}

func CORSOptions() {
	handlers.AllowedOrigins([]string{"*"})
	handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Host"})
	handlers.AllowedMethods([]string{"POST", "GET"})
}
