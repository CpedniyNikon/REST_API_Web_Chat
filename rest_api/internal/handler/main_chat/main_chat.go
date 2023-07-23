package main_chat

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"rest_api/internal/handler/utils"
	"strconv"
	time "time"
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

type UserInfo struct {
	UserID int `json:"userID"`
}

func (h *Handler) getMessages(c *gin.Context) {

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

	var UserInfo UserInfo
	if err := c.ShouldBindJSON(&UserInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var UserID = UserInfo.UserID

	var Login string
	var Password string
	var IsLogged bool
	var YourTimeLogged time.Time
	var TimeLogged time.Time

	query :=
		"select * from userdata where id=$1"
	err = db.QueryRow(query, UserID).Scan(&UserID, &Login, &Password, &IsLogged, &YourTimeLogged)
	if err != nil {
		log.Fatal(err)
	}

	query =
		"select * from message"
	rows, _ := db.Query(query)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	CurrentTime := time.Now()
	fmt.Println(CurrentTime)

	var messages []utils.MessageResponse
	for rows.Next() {
		var Id int
		var Message string
		var TimeSended time.Time
		var UserId int
		err := rows.Scan(&Id, &Message, &TimeSended, &UserId)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("current time:= ")
		fmt.Println(CurrentTime)

		fmt.Print("time sended:= ")
		fmt.Println(TimeSended)

		query =
			"select * from userdata where id=$1"
		err = db.QueryRow(query, UserId).Scan(&Id, &Login, &Password, &IsLogged, &TimeLogged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Id := " + strconv.Itoa(UserId))
		fmt.Println("login := " + Login)

		if YourTimeLogged.After(TimeSended) {
			continue
		}

		messages = append(messages, utils.MessageResponse{User: Login, Message: Message})
	}

	c.JSON(http.StatusOK, messages)
}
