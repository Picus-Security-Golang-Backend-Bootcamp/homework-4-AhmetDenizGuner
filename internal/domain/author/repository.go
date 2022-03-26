package author

import "gorm.io/gorm"

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) InsertInitialData(authors []Author) {

	for _, author := range authors {
		r.db.Create(&author)
	}
}

func (r *AuthorRepository) FindByName(name string) (Author, error) {
	var author Author
	result := r.db.Where("Name = ?", name).Find(&author)
	if result.Error != nil {
		return Author{}, result.Error
	}
	return author, nil
}

func (r *AuthorRepository) FindByID(id int) (Author, error) {
	var author Author
	result := r.db.First(&author, id)
	if result.Error != nil {
		return Author{}, result.Error
	}
	return author, nil
}

func (r *AuthorRepository) Create(a Author) error {
	result := r.db.Create(a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
