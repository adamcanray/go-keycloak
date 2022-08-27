package domains

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	ID          int    `json:"id" gorm:"unique"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
