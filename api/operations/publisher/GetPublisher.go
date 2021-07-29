package publisher

import (
	"new_library/env"
	"new_library/model"
)

func GetPublisher() ([]*model.Publishers, error) {
	db := env.RDB
	publisher := []*model.Publishers{}
	db.Find(&publisher)

	return publisher, nil

}
