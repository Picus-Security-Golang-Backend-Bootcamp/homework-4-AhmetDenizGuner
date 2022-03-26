package book

import "gorm.io/gorm"

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) FindAll() ([]Book, error) {
	var books []Book

	result := r.db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (r *BookRepository) InsertInitialData(books []Book) {

	for _, book := range books {
		r.db.Create(&book)
	}
}

//search book that contains searched key in stock code, isbn and book name column
func (r *BookRepository) FindAllByKey(key string) ([]Book, error) {
	var books []Book
	key = "%" + key + "%"
	result := r.db.Preload("Author").Joins("join table_author athr on athr.id = table_book.author_id").Where("table_book.Name ILIKE ? OR table_book.ISBN ILIKE ? OR table_book.Stock_Code ILIKE ? OR athr.Name ILIKE ?", key, key, key, key).Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (r *BookRepository) DeleteById(id int) error {
	result := r.db.Delete(&Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *BookRepository) FindById(id int) (Book, error) {
	var book Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return Book{}, result.Error
	}
	return book, nil
}

func (r *BookRepository) Update(b Book) error {
	result := r.db.Save(b)
	//r.db.Model(&c).Update("name", "deneme")

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *BookRepository) Create(b Book) error {
	result := r.db.Create(b)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *BookRepository) FindByName(name string) (Book, error) {
	var book Book
	result := r.db.Where("Name = ?", name).Find(&book)
	if result.Error != nil {
		return Book{}, result.Error
	}
	return book, nil
}
