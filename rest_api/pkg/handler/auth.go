package handler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

type UserData struct {
	Login    string
	Password string
}

type LogInResponse struct {
	Text string
}

type UserAlreadyExistsError struct {
	message string
}

func (err UserAlreadyExistsError) Error() string {
	return err.message
}

func InsertUser(user UserData, db *sql.DB) error {
	fmt.Println(user.Login)
	fmt.Println(user.Password)

	var exist bool
	query := "SELECT EXISTS(SELECT login, password FROM userdata WHERE login=$1 AND password=$2)"
	err := db.QueryRow(query, user.Login, user.Password).Scan(&exist)
	if err != nil {
		panic(err)
	}

	if exist {
		return UserAlreadyExistsError{"Пользователь уже существует в базе данных."}
	}

	fmt.Println("inserting")
	_, _ = db.Exec("insert into \"userdata\" (login, password) values ($1, $2)",
		user.Login, user.Password)
	return nil
}

type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (h *Handler) signUp(c *gin.Context) {
	fmt.Println("sing-up")

	var user UserData

	if err := c.ShouldBindJSON(&user); err != nil {
		//Ошибка привязки JSON
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn := connection{
		Host:     "postgres",
		Port:     "5432",
		User:     "postgres",
		Password: "qwerty",
		DBName:   "postgres",
	}

	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conn.Host,
		conn.Port,
		conn.User,
		conn.Password,
		conn.DBName,
	)
	fmt.Println(connInfo)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()
	err = InsertUser(user, db)
	if err != nil {
		response := LogInResponse{"user already exists"}
		c.JSON(http.StatusOK, response)
	} else {
		response := LogInResponse{"new user added to db"}
		c.JSON(http.StatusOK, response)
	}
}

func (h *Handler) signIn(c *gin.Context) {
	var user UserData

	if err := c.ShouldBindJSON(&user); err != nil {
		//Ошибка привязки JSON
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn := connection{
		Host:     "postgres",
		Port:     "5432",
		User:     "postgres",
		Password: "qwerty",
		DBName:   "postgres",
	}

	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conn.Host,
		conn.Port,
		conn.User,
		conn.Password,
		conn.DBName,
	)
	fmt.Println(connInfo)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()

	var exist bool
	query := "SELECT EXISTS(SELECT login, password FROM userdata WHERE login=$1 AND password=$2)"
	err = db.QueryRow(query, user.Login, user.Password).Scan(&exist)
	if err != nil {
		panic(err)
	}

	if exist {
		response := LogInResponse{"user already exists"}
		c.JSON(http.StatusOK, response)
	} else {
		response := LogInResponse{"no such user id db"}
		c.JSON(http.StatusOK, response)
	}
}
