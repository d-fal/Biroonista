package routes

import (
	"net/http"

	"../controllers/index"
	"../controllers/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func InitRoutes(router *gin.Engine) {

	router.Use(cors.Default())
	router.LoadHTMLGlob("./../../../client/front/*.html")
	router.GET("/", index.EnSite)
	router.Static("/static", "./../../../client/front/static")
	router.Use(favicon.New("./../../../client/front/favicon.ico"))

	var authAgent users.AuthMiddleware
	var registerAgent users.NewUser
	var loginAgent users.Login

	registeration := router.Group("/register")

	registeration.Use(registerAgent.RegisterationMiddleware())
	{
		registeration.POST("/captcha", registerAgent.CaptchaHandler)
		registeration.POST("/", registerAgent.RegisterationHandler)
		registeration.PUT("/acaptcha", registerAgent.GetAudibleCaptcha)
	}
	loginRoute := router.Group("/login")
	router.GET("/login", index.EnSite)
	router.GET("/register", index.EnSite)
	router.GET("/login/forget", index.EnSite)
	router.GET("/user", index.EnSite)
	loginRoute.Use(loginAgent.LoginMiddleware())
	{

		loginRoute.POST("/", loginAgent.LoginHandler)
		loginRoute.POST("/forget", loginAgent.ForgetPasswordHandler)
	}

	userRoute := router.Group("/user")
	userRoute.Use(authAgent.AuthenticateMiddleware())
	{
		userRoute.POST("/renew_credentials", authAgent.RefreshTokenHandler)
		userRoute.POST("/basicinfo", authAgent.MarshalBasicInfoHandler)
		userRoute.POST("/sendtext", authAgent.SendVerificationCodeHandler)
		userRoute.POST("/verifycontact", authAgent.ContactVerificationByCodeHandler)
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router.Use(static.Serve("/", static.LocalFile("../../../../client/front", true)))
	router.Use(static.Serve("/repo", static.LocalFile("./captcha", true)))

}
