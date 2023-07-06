package utils

type UserData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type MessageData struct {
	Message string `json:"message"`
	UserId  int    `json:"userId"`
}
