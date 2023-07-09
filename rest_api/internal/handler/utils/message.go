package utils

type MessageData struct {
	Message string `json:"message"`
	UserId  int    `json:"userId"`
}

type MessageResponse struct {
	User    string `json:"user"`
	Message string `json:"message"`
}
