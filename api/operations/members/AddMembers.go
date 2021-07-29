package members

import (
	"new_library/env"
	"new_library/model"
)

func AddBooks(members *model.Members) (*model.Members, error) {
	db := env.RDB
	err := db.Create(&members).Error
	if err != nil {
		return nil, err
	} else {
		return members, nil
	}
}
