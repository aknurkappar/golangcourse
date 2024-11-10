package api

import (
	_ "github.com/go-playground/validator/v10"
	"time"
)

type Role struct {
	ID   uint   `json:"id" gorm:"primary_key" xml:"id"`
	Name string `json:"name" gorm:"unique" xml:"name"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primary_key" xml:"id"`
	Name     string `json:"name" xml:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" gorm:"unique" xml:"email" binding:"required,email"`
	Password string `json:"password" xml:"password" binding:"required,min=8"`
	RoleID   uint   `json:"role_id" gorm:"default:2" xml:"role_id"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleID" xml:"role"`
}

type Post struct {
	ID        uint      `json:"id" gorm:"primary_key" xml:"id"`
	UserID    uint      `json:"user_id" xml:"user_id" binding:"required"`
	Text      string    `json:"text" xml:"text" binding:"required,min=1,max=500"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" xml:"created_at"`
}
