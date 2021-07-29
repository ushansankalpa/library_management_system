package publisher

import (
	"new_library/env"
	"new_library/model"
)

func AddPublisher(publisher *model.Publishers) (*model.Publishers, error) {
	db := env.RDB

	err := db.Create(&publisher).Error

	if err != nil {
		return nil, err
	} else {
		return publisher, nil
	}

}
