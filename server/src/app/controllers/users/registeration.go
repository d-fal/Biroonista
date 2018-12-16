package users

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/dchest/captcha"

	"../../config/db"
	"../../lib"
	"../../model"
	"github.com/gin-gonic/gin"
)

type NewUser struct {
	Password        string `form:"password" json:"password" binding:"required"`
	Retype_Password string `form:"retype_password" json:"retype_password" binding:"required"`
	Tel             string `form:"tel" json:"tel" binding:"required"`
	Captcha         string `form:"captcha" json:"captcha" binding:"required"`
	CaptchaStr      string `form:"captcha_str" json:"captcha_str" binding:"required"`
	ACaptcha        string `form:"acaptcha" json:"acaptcha" binding:"required"`
	Secret          []byte
	Expiration      int
}

func (newUser *NewUser) RegisterationHandler(c *gin.Context) {
	var userInstance model.Users
	var token string
	var exp time.Time
	if err := c.ShouldBind(&newUser); err != nil {
		panic("Cannot bind")
	}
	errorCode := newUser.CheckRegisterationInput()

	if errorCode == lib.NO_ERROR {
		pgdb := db.DbModel()
		defer pgdb.Close()

		userInstance.Username = newUser.Tel
		userInstance.Password = newUser.Password
		_, err := pgdb.Model(&userInstance).Value("password", "crypt(?password, 'bf')", newUser.Password).Insert()

		if err != nil {
			errorCode = lib.WARNING_USER_ALREADY_EXISTS
			panic(err)
		}

		usersContactInstance := model.UsersContacts{
			Users:        &userInstance,
			UsersId:      userInstance.Id,
			Address:      newUser.Tel,
			IsMain:       true,
			IsVerified:   false,
			Verifier:     fmt.Sprintf("%d", 100000+rand.Intn(899999)),
			ContactsType: 1,
		}
		if err == nil {
			if _, err = pgdb.Model(&usersContactInstance).Insert(); err != nil {

				errorCode = lib.WARNING_CANNOT_SAVE_CONTACT
				panic(err)
			}

			// Generate token
			token, exp = GenerateToken(userInstance.Username, lib.Secret)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"response": lib.ErrorMessages()[errorCode],
		"error":    errorCode,
		"token":    token,
		"exp":      exp,
	})
}

func (newUser *NewUser) CheckRegisterationInput() int {

	errorCode := lib.NO_ERROR

	newUser.Tel, errorCode = SanitizeTel(newUser.Tel)

	if newUser.Password != newUser.Retype_Password {
		errorCode = lib.WARNING_PASSWORDS_DOES_NOT_MATCH
	}
	if len(newUser.Password) < 6 {
		errorCode = lib.WARNING_PASSWORD_TOO_SHORT

	}

	if !captcha.VerifyString(newUser.Captcha, newUser.CaptchaStr) {
		errorCode = lib.WARNING_INCORRECT_CAPTCHA
	}

	return errorCode

}

func (newUser *NewUser) RegisterationMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

	}
}
