package models

type Todo struct {
	//gorm.Model
	TodoID      uint       `json:"todoid" gorm:"primary_key"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Catgories   []Category `json:"catgories" gorm:"foreignkey:TodoID"`
}
