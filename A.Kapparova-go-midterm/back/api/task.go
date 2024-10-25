package api

import (
	"time"
)

type Task struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	Category   *Category `json:"category" gorm:"foreignKey:CategoryID"`
	CategoryID *uint     `json:"category_id,omitempty"`
}

type Category struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name" gorm:"unique"`
}
