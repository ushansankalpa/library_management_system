package model

type Books struct {
	Model
	BookId    int    `json:"book_id" swaggerignore:"true"`
	BookTitle string `json:"book_title"`
	Author    string `json:"author"`
	Price     int    `json:"price"`
	Available string `json:"available"`
}
