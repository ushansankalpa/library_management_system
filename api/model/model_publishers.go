package model

type Publishers struct {
	Model
	PubId      int    `json:"pub_id" swaggerignore:"true"`
	PubName    string `json:"pub_name"`
	PubAddress string `json:"pub_address"`
	Books      *Books `gorm:"foreignkey:book_id" json:"book_id"`
}
