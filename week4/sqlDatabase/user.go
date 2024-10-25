package sqlDatabase

type UserForm struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}
