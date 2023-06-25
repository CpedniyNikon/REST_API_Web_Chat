package handler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserData struct {
	Id       int
	Login    string
	Password string
}

func InsertUser(user UserData, db *sql.DB) {
	var exist bool
	query := "SELECT EXISTS(SELECT login, password FROM userdata WHERE login=$1 AND password=$2)"
	err := db.QueryRow(query, user.Login, user.Password).Scan(&exist)
	if err != nil {
		panic(err)
	}

	if !exist {
		_, _ = db.Exec("insert into \"userdata\" (login, password) values ($1, $2)",
			user.Login, user.Password)
	}
}

func (h *Handler) signUp(c *gin.Context) {
	fmt.Println("sing-up")
	var user UserData

	if err := c.ShouldBindJSON(&user); err != nil {
		//Ошибка привязки JSON
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, err := sql.Open("postgres", "user=postgres password=04122002 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	InsertUser(user, db)
}

func (h *Handler) signIn(c *gin.Context) {

}
