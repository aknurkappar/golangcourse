package api

type Role struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"unique"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id" gorm:"default:2"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleID"`
}
