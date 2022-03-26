package author

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
}

//NewAuthor is struct constructor of Author model
func NewAuthor(authorName string) *Author {
	author := &Author{
		Name: authorName,
	}
	return author
}
