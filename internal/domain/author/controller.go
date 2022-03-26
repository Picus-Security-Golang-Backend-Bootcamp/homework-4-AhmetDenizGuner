package author

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	responsemodel "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/pkg/response_model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AuthorHandler struct {
	db               *gorm.DB
	authorRepository AuthorRepository
}

func NewAuthorHandler(db *gorm.DB) *AuthorHandler {
	return &AuthorHandler{
		db:               db,
		authorRepository: *NewAuthorRepository(db),
	}
}

func (ah *AuthorHandler) GetAuthorByIdHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	author, err := ah.authorRepository.FindByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	d := responsemodel.ApiResponse{
		Data: author,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp, _ := json.Marshal(d)
	w.Write(resp)

}

func (ah *AuthorHandler) GetAuthorByNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]

	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(r.RequestURI + ": missing arguments")
		return
	}

	author, err := ah.authorRepository.FindByName(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	d := responsemodel.ApiResponse{
		Data: author,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp, _ := json.Marshal(d)
	w.Write(resp)

}

func (ah *AuthorHandler) AddAuthorHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name, ok := vars["name"]

	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(r.RequestURI + ": missing arguments")
		return
	}

	author, err := ah.authorRepository.FindByName(name)

	if err == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(r.RequestURI + ": Author is already exist")
		return
	}

	author.Name = name

	err = ah.authorRepository.Create(author)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
	}

	d := responsemodel.ApiResponse{
		Data: author.Name + " added",
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp, _ := json.Marshal(d)
	w.Write(resp)
}
