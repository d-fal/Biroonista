package main

import (
	"net/http"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"./config/mydb"

)


func initRoutes(router *gin.Engine) {
	// Setup homepage route
	
	router.GET("/",enSite)
	router.GET("/login", loginPanel)
	// static files
	router.Use(static.Serve("/", static.LocalFile("../../../client/", true)))
	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
}
func main() {
	// Default Gin router
	mydb.PgORMTest()
	mydb.CreateModels()
	// // mydb.ExampleInsert()
	mydb.CloseDb()
	
	router := gin.Default()
	
	router.LoadHTMLGlob("../../../client/views/*/*.html")
	

	initRoutes(router)
	router.Run(":5000")
}

func enSite(c *gin.Context){
	// c.Header("content-type" , "application/html")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"body":"Biroonista",
		"title":"Biroonista App",
	})
}

func loginPanel(c *gin.Context) {
	c.HTML( http.StatusOK, "login.html", gin.H{
		"body":"no",
	})
}

