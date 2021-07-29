package model

type Members struct {
	Model
	MemberId      int    `json:"mem_id"`
	MemberName    string `json:"mem_name"`
	MemberAddress string `json:"mem_address"`
	MemberType    string `json:"mem_type"`
	MemberDate    string `json:"mem_date"`
	ExpDate       string `json:"exp_date"`
	Books         *Books `gorm:"foreignkey:book_id" json:"book_id"`
}
