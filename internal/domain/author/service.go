package author

import (
	"errors"
	"log"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/database"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/pkg/csv_helper"
)

func GetAuthorByID(id int, authorRepository *AuthorRepository) (Author, error) {

	author, err := authorRepository.FindByID(id)

	if err != nil {
		return Author{}, err
	}

	return author, nil

}

func InsertInitialAuthorData(authorRepository *AuthorRepository) error {

	dbExist := database.DB.Migrator().HasTable(&Author{})

	if !dbExist {
		database.DB.AutoMigrate(&Author{})

		// read CSV
		authorSlice, err := csv_helper.ReadCsv("../resources/author.csv", 1)

		if err != nil {
			return errors.New("CSV cannot be read!")
		}

		log.Println(authorSlice)

		//create author data
		var authorData []Author
		for _, author := range authorSlice {

			newAuthor := NewAuthor(author[1])
			authorData = append(authorData, *newAuthor)

		}

		// add author data
		authorRepository.InsertInitialData(authorData)

	}

	return nil

}
