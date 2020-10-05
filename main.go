package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
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

func indexR(c *gin.Context) {
	ans, ok := c.Get("test")
	if !ok {
		ans = false
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "index",
		"test":    ans,
	})
}

func login(c *gin.Context) {
	author := c.Request.Header.Get("Authorization")
	fmt.Println(author)
	if author == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "401",
			"message": "登录信息不正确",
			"data":    "",
		})
		c.Abort()
		return
	}
	auth := ""
	ErrAuth := db.QueryRow("select auth from authorization where auth=?;", author).Scan(&auth)
	if ErrAuth != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "后台程序错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	if auth == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "401",
			"message": "用户登录信息错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	c.Next()
}

func main() {
	router := gin.Default()
	db, _ = sql.Open("mysql", "root:yangxuechao123@/medicalsystem?charset=utf8")
	defer db.Close()
	api := router.Group("/api")
	{
		api.GET("/ping", login, func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		api.GET("/dataset", login, func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		api.POST("/login/", login, func(c *gin.Context) {
			username := c.PostForm("username")
			pwdCrypto := sha256.Sum256([]byte(c.PostForm("password")))
			password := hex.EncodeToString(pwdCrypto[:])
			var pwd string
			err1 := db.QueryRow("select password from user where username=?;", username).Scan(&pwd)
			if err1 != nil {
				fmt.Printf("%#v 55\n", err1)
			}
			if pwd == "" {
				c.JSON(200, gin.H{
					"err": "账号不存在",
				})
			} else if pwd != password {
				fmt.Println(pwd)
				fmt.Println(password)
				c.JSON(200, gin.H{
					"err": "密码错误",
				})
			} else {
				c.JSON(200, gin.H{
					"data": gin.H{
						"access": "thisistokentest",
					},
				})
			}
		})

		api.GET("/api/dataset/", func(c *gin.Context) {
			username := c.PostForm("username")
			pwdCrypto := sha256.Sum256([]byte(c.PostForm("password")))
			password := hex.EncodeToString(pwdCrypto[:])
			var pwd string
			err1 := db.QueryRow("select password from user where username=?;", username).Scan(&pwd)
			if err1 != nil {
				fmt.Printf("%#v 55\n", err1)
			}
			if pwd == "" {
				c.JSON(200, gin.H{
					"err": "账号不存在",
				})
			} else if pwd != password {
				fmt.Println(pwd)
				fmt.Println(password)
				c.JSON(200, gin.H{
					"err": "密码错误",
				})
			} else {
				c.JSON(200, gin.H{
					"data": gin.H{
						"access": "thisistokentest",
					},
				})
			}
		})

	}
	router.Run(":8080")
}
