package handler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"rest_api/pkg/handler/utils"
	"rest_api/pkg/handler/utils/errors"
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

func InsertUser(user utils.UserData, db *sql.DB) error {
	fmt.Println(user.Login)
	fmt.Println(user.Password)

	var exist bool
	query := "SELECT EXISTS(SELECT login, password FROM userdata WHERE login=$1 AND password=$2)"
	err := db.QueryRow(query, user.Login, user.Password).Scan(&exist)
	if err != nil {
		panic(err)
	}

	if exist {
		return errors.UserAlreadyExistsError{Message: "Пользователь уже существует в базе данных."}
	}

	fmt.Println("inserting")
	_, _ = db.Exec("insert into \"userdata\" (login, password) values ($1, $2)",
		user.Login, user.Password)
	return nil
}

func (h *Handler) signUp(c *gin.Context) {
	var user utils.UserData
	if err := c.ShouldBindJSON(&user); err != nil {
		//Ошибка привязки JSON
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
	err = InsertUser(user, db)
	if err != nil {
		response := utils.RequestResponse{Text: "user already exists"}
		c.JSON(http.StatusOK, response)
	} else {
		response := utils.RequestResponse{Text: "new user added to db"}
		c.JSON(http.StatusOK, response)
	}
}

func (h *Handler) signIn(c *gin.Context) {
	var user utils.UserData
	if err := c.ShouldBindJSON(&user); err != nil {
		//Ошибка привязки JSON
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

	var exist bool
	query := "SELECT EXISTS(SELECT login, password FROM userdata WHERE login=$1 AND password=$2)"
	err = db.QueryRow(query, user.Login, user.Password).Scan(&exist)
	if err != nil {
		panic(err)
	}

	if exist {

		query =
			`UPDATE postgres.public.userdata
		SET time_logged = $1, is_logged = $2
		WHERE login = $3 and password = $4
		`

		login, password, logged, loginTime := user.Login, user.Password, 1, time.Now()

		_, err = db.Exec(query, loginTime, logged, login, password)
		if err != nil {
			panic(err)
		}
		response := utils.RequestResponse{Text: "u just logged in"}
		c.JSON(http.StatusOK, response)
	} else {
		response := utils.RequestResponse{Text: "no such user id db"}
		c.JSON(http.StatusOK, response)
	}
}

func (h *Handler) signOut(c *gin.Context) {
	var user utils.UserData
	if err := c.ShouldBindJSON(&user); err != nil {
		//Ошибка привязки JSON
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

	query :=
		`UPDATE postgres.public.userdata
		SET time_logged = $1, is_logged = $2
		WHERE login = $3 and password = $4
		`

	login, password, logged := user.Login, user.Password, 0

	var loginTime *time.Time = nil

	_, err = db.Exec(query, loginTime, logged, login, password)
	if err != nil {
		panic(err)
	}
}
