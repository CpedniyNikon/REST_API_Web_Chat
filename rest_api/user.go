package rest_api

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
