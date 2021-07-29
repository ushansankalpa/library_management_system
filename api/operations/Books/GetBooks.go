package books

import (
	"new_library/env"
	"new_library/model"
)

func GetBooks() ([]*model.Books, error) {
	db := env.RDB
	books := []*model.Books{}
	db.Find(&books)
	return books, nil
}
