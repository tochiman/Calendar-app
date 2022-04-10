package main

import (
	"crypto"
	"log"
	"net/http"

	"./service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/src", "./src")
	router.LoadHTMLGlob("html/*")

	router.GET("/",func (c *gin.Context){
		c.HTML(http.StatusOK,"app.html",gin.H{
		})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "login-page",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		service.DbInit()
		dbpasswd := service.DbGetOne(c.PostForm("login")).Password
		log.Print(dbpasswd)
		formpasswd := c.PostForm("password")
		log.Print(formpasswd)

		if err := crypto.CompareHashAndPassword(dbpasswd, formpasswd); err != nil {
            log.Println("ログインできませんでした")
            c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": "ログインすることができませんでした"})
            c.Abort()
        } else {
            log.Println("ログインできました")
            c.Redirect(302, "/")
		}
	})

	router.Run("0.0.0.0:58080")
}