package main_chat

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"rest_api/internal/handler/utils"
	"time"
)

var connection = utils.DBConnection{
	Host:     "postgres",
	Port:     "5432",
	User:     "postgres",
	Password: "qwerty",
	DBName:   "postgres",
}
var connectionInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	connection.Host,
	connection.Port,
	connection.User,
	connection.Password,
	connection.DBName)

func AddMessage(message utils.MessageData, db *sql.DB) {
	fmt.Println(message.Message)
	fmt.Println(message.UserId)

	fmt.Println("adding message")

	query :=
		"insert into \"message\" (text_message, time_sended, user_id) values ($1, $2, $3)"
	_, _ = db.Exec(query, message.Message, time.Now(),
		message.UserId)
}

func (h *Handler) write(c *gin.Context) {

	var user utils.MessageData
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	AddMessage(user, db)
}
