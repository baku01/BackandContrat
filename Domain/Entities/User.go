package Entities

type User struct {
	Id       string
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
