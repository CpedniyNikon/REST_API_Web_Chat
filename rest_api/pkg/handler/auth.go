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

func InsertUser(user UserData, db *sql.DB) {
	fmt.Println(user.Login)
	fmt.Println(user.Password)

	var exist bool
	exist = false
	query := "SELECT EXISTS(SELECT login, password FROM userdata WHERE login=$1 AND password=$2)"
	err := db.QueryRow(query, user.Login, user.Password).Scan(&exist)
	if err != nil {
		panic(err)
	}

	if !exist {
		fmt.Println("inserting")
		_, _ = db.Exec("insert into \"userdata\" (login, password) values ($1, $2)",
			user.Login, user.Password)
	}
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
	InsertUser(user, db)
}

func (h *Handler) signIn(c *gin.Context) {

}
