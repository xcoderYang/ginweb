package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// type db struct{

// }
// func (db *db) insert(tables, cols, values){

// }

type User struct {
	id         string
	username   string
	password   string
	add_time   string
	power      string
	nickname   string
	label_id   string
	mobile     string
	father     string
	email      string
	last_login string
}

func m1(c *gin.Context) {
	fmt.Println("m1 in")
	c.Next()
	fmt.Println("m1 out")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in")
	c.Set("test", true)
	fmt.Println("m2 out")
}

func indexR(c *gin.Context) {
	fmt.Println("index")
	ans, ok := c.Get("test")
	if !ok {
		ans = false
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "index",
		"test":    ans,
	})
}

func main() {
	router := gin.Default()
	router.Use(m1, m2, indexR)
	router.GET("/index", func(c *gin.Context) {
	})
	// api := router.Group("/api")
	// fmt.Println("one")
	// db, err := sql.Open("mysql", "root:yangxuechao123@/medicalsystem?charset=utf8")
	// {
	// 	api.GET("/ping", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 			"message": "pong",
	// 		})
	// 	})
	// 	api.POST("/login/", func(c *gin.Context) {
	// 		db, err := sql.Open("mysql", "root:yangxuechao123@/medicalsystem?charset=utf8")
	// 		if err != nil {
	// 			fmt.Println(err, "46")
	// 		}
	// 		defer db.Close()
	// 		username := c.PostForm("username")
	// 		pwdCrypto := sha256.Sum256([]byte(c.PostForm("password")))
	// 		password := hex.EncodeToString(pwdCrypto[:])
	// 		var pwd string
	// 		err1 := db.QueryRow("select password from user where username=?;", username).Scan(&pwd)
	// 		if err1 != nil {
	// 			fmt.Printf("%#v 55\n", err1)
	// 		}
	// 		if pwd == "" {
	// 			c.JSON(200, gin.H{
	// 				"err": "账号不存在",
	// 			})
	// 		} else if pwd != password {
	// 			fmt.Println(pwd)
	// 			fmt.Println(password)
	// 			c.JSON(200, gin.H{
	// 				"err": "密码错误",
	// 			})
	// 		} else {
	// 			c.JSON(200, gin.H{
	// 				"data": gin.H{
	// 					"access": "thisistokentest",
	// 				},
	// 			})
	// 		}
	// 	})

	// 	api.GET("/api/dataset/", func(c *gin.Context) {
	// 		db, err := sql.Open("mysql", "root:yangxuechao123@/medicalsystem?charset=utf8")
	// 		if err != nil {
	// 			fmt.Println(err, "46")
	// 		}
	// 		defer db.Close()
	// 		username := c.PostForm("username")
	// 		pwdCrypto := sha256.Sum256([]byte(c.PostForm("password")))
	// 		password := hex.EncodeToString(pwdCrypto[:])
	// 		var pwd string
	// 		err1 := db.QueryRow("select password from user where username=?;", username).Scan(&pwd)
	// 		if err1 != nil {
	// 			fmt.Printf("%#v 55\n", err1)
	// 		}
	// 		if pwd == "" {
	// 			c.JSON(200, gin.H{
	// 				"err": "账号不存在",
	// 			})
	// 		} else if pwd != password {
	// 			fmt.Println(pwd)
	// 			fmt.Println(password)
	// 			c.JSON(200, gin.H{
	// 				"err": "密码错误",
	// 			})
	// 		} else {
	// 			c.JSON(200, gin.H{
	// 				"data": gin.H{
	// 					"access": "thisistokentest",
	// 				},
	// 			})
	// 		}
	// 	})

	// }
	router.Run(":8080")
}
