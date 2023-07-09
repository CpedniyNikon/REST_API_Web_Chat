package auth

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"rest_api/internal/handler/utils"
	"rest_api/internal/handler/utils/errors"
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

func AddUser(user utils.UserData, db *sql.DB) error {
	fmt.Println(user.Login)
	fmt.Println(user.Password)

	var exist bool
	query := "select exists(select login, password from userdata where login=$1 and password=$2)"
	err := db.QueryRow(query, user.Login, user.Password).Scan(&exist)
	if err != nil {
		panic(err)
	}

	if exist {
		return errors.UserAlreadyExistsError{Message: "user already exists in db"}
	}

	fmt.Println("adding user")
	_, _ = db.Exec("insert into \"userdata\" (login, password) values ($1, $2)",
		user.Login, user.Password)
	return nil
}

func (h *Handler) signUp(c *gin.Context) {
	var user utils.UserData
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
	err = AddUser(user, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "user already exists"})
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

func (h *Handler) signIn(c *gin.Context) {
	var user utils.UserData
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

	var exist bool
	query := "select exists(select login, password from userdata where login=$1 and password=$2)"
	err = db.QueryRow(query, user.Login, user.Password).Scan(&exist)
	if err != nil {
		panic(err)
	}

	if exist {

		query =
			`update postgres.public.userdata
		set time_logged = $1, is_logged = $2
		where login = $3 and password = $4 returning id
		`

		login, password, logged, loginTime := user.Login, user.Password, 1, time.Now()

		var id int64
		_ = db.QueryRow(query, loginTime, logged, login, password).Scan(&id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"id": id})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"response": "there is no user with such parameters"})
	}
}

func (h *Handler) signOut(c *gin.Context) {
	var user utils.UserData
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

	query :=
		`update postgres.public.userdata
		set time_logged = $1, is_logged = $2
		where login = $3 and password = $4
		`

	login, password, logged := user.Login, user.Password, 0

	var loginTime *time.Time = nil

	_, err = db.Exec(query, loginTime, logged, login, password)
	if err != nil {
		panic(err)
	}
}
