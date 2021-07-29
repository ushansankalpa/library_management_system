package members

import (
	"new_library/env"
	"new_library/model"
)

func GetMember() ([]*model.Members, error) {
	db := env.RDB

	members := []*model.Members{}
	db.Find(&members)

	return members, nil
}
