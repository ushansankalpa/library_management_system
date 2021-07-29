package books

import (
	"new_library/env"
	"new_library/model"
)

func AddBooks(books *model.Books) (*model.Books, error) {
	db := env.RDB
	err := db.Create(books).Error
	if err != nil {
		return nil, err
	} else {
		return books, nil
	}
}
