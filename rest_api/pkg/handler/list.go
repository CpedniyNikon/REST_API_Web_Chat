package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func (h *Handler) createList(c *gin.Context) {
	fmt.Println("post api lists")
}

//var db* sql.DB
//
//func DoShit()  {
//	var err error
//	db, err = sql.Open("postgres", "user=postgres password=04122002 dbname=postgres sslmode=disable")
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//
//	rows, err := db.Query("SELECT * FROM \"User_data\"")
//	if err != nil{
//		panic(err)
//	}
//	defer rows.Close()
//
//	var users []UserData
//
//	for rows.Next(){
//		u := UserData{}
//		err := rows.Scan(&u.Email, &u.Password, &u.Username, &u.Id)
//		if err != nil{
//			fmt.Println(err)
//			continue
//		}
//		users = append(users, u)
//	}
//	for _, u := range users{
//		fmt.Println(u.Email, u.Password, u.Username, u.Id)
//	}
//
//}

func (h *Handler) getAllLists(c *gin.Context) {
	fmt.Println("get api lists")

}
func (h *Handler) getListById(c *gin.Context) {

}
func (h *Handler) updateList(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
