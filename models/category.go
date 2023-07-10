package models

type Category struct {
	CatgoryID uint   `json:"catgoryid" gorm:"primary_key"`
	Name      string `json:"name"`
	TodoID    uint   `json:"-"`
}
