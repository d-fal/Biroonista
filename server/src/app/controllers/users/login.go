package users

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"../../config/db"
	"../../lib"
	"../../model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	kavenegar "github.com/kavenegar/kavenegar-go"
)

type VerifyUser struct {
	Verification string `form:"verify" json:"verify" binding:"required"`
}

type Login struct {
	Username      string `form:"username" json:"username" binding:"required"`
	Password      string `form:"password" json:"password" binding:"required"`
	Token         string
	Authorization string
}
type ResetPass struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Pattern  string `json:"pattern"`
	TypeOf   int    `json:"type_of"`
	Bearer   string
	Token    string
}

func (l *Login) Authenticate(c *gin.Context, credentials ResetPass) (bool, jwt.MapClaims) {
	bearer := c.GetHeader("Authorization")
	credentials.Bearer = bearer
	var token string
	slices := strings.Split(bearer, " ")

	if slices[0] == "Bearer" {
		if len(slices) == 2 {
			token = slices[1]
			l.Token = token
		}
	}

	return l.VerifyToken(token, credentials)
}

func (l *Login) VerifyToken(tokenString string, credentials ResetPass) (bool, jwt.MapClaims) {
	if tokenString == "" {
		return false, nil
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, nil
		}
		return lib.SecretTemp, nil
	})

	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		credentials.Token = tokenString
		return true, claim
	} else {
		return false, claim
	}

}
func (l *Login) LoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func (l *Login) LoginHandler(c *gin.Context) {
	var login Login
	errorCode := lib.NO_ERROR
	if err := c.ShouldBind(&login); err != nil {
		panic("Cannot bind")
	}
	var user ClaimedUser
	user.UserName, errorCode = SanitizeTel(login.Username)
	user.Password = login.Password

	token, validUntil := GenerateToken(user.UserName, lib.Secret)

	if isAuth, isVerified := AuthUser(user); isAuth {

		errorCode = lib.LOGIN_SUCCESSFUL
		if !isVerified {
			errorCode = lib.WARNING_CONTACT_NOT_VERIFIED
		}
		c.JSON(http.StatusOK, gin.H{
			"token":    token,
			"exp":      validUntil,
			"error":    errorCode,
			"response": lib.ErrorMessages()[errorCode],
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error":    lib.WARNING_LOGIN_FAILED_INCORRECT_CREDENTIALS,
			"response": lib.ErrorMessages()[lib.WARNING_LOGIN_FAILED_INCORRECT_CREDENTIALS],
		})
	}

}

func (l *Login) ForgetPasswordHandler(c *gin.Context) {
	var (
		token string
	)
	errorCode := lib.NO_ERROR
	var credentials ResetPass
	var usersContact model.UsersContacts
	if err := c.ShouldBind(&credentials); err != nil {
		errorCode = lib.ERROR_BINDING_FAILED
	}
	if credentials.TypeOf == lib.TYPE_CONTACT_MOBILE {
		credentials.Username, errorCode = SanitizeTel(credentials.Username)

	}
	usersContact, errorCode = GetUserByContact(credentials.Username)
	if errorCode < 0 {
		if isAuth, claim := l.Authenticate(c, credentials); !isAuth {
			token, _ = GenerateToken(credentials.Username, lib.SecretTemp)
			errorCode = lib.RESET_TOKEN_GENERATED
			//send Text

			SendResetPaaswordText(usersContact, credentials.Username)
		} else {
			if claim["user"].(string) != credentials.Username {
				errorCode = lib.WARNING_SUSCPICIOUS_USER
			}
			if usersContact.Verifier == credentials.Pattern {
				errorCode = lib.PATTERN_ACCEPTED
				token = l.Token
				if credentials.Password != "" {

					if len(credentials.Password) < 6 {
						errorCode = lib.WARNING_PASSWORD_TOO_SHORT

					} else {
						errorCode = credentials.ChangePassword(usersContact)
					}

				}

			} else {
				errorCode = lib.WARNING_INVALID_VERIFICATION_PATTERN
			}

		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error":        errorCode,
		"response":     lib.ErrorMessages()[errorCode],
		"client_token": token,
	})
}
func (credentials *ResetPass) ChangePassword(usersContact model.UsersContacts) int {
	var users model.Users
	errorCode := lib.PASSWORD_CHANGED
	pgdb := db.DbModel()
	defer pgdb.Close()
	_, err := pgdb.Model(&users).Set("password =crypt(?, 'bf') ",
		credentials.Password).Where("id =?", usersContact.UsersId).Update()
	if err != nil {
		errorCode = lib.WARNING_CANNOT_CHANGE_PASSWORD
	}
	return errorCode
}
func SendResetPaaswordText(usersContact model.UsersContacts, user string) int {
	errorCode := lib.NO_ERROR

	diff := time.Now().Sub(usersContact.VerificationDate)
	pgdb := db.DbModel()
	defer pgdb.Close()
	usersContact.Verifier = fmt.Sprintf("%d", 100000+rand.Intn(899999))

	if int(diff.Seconds()) < lib.VERIFICATION_EXPIRATION_SECONDS {
		errorCode = lib.WARNING_LAST_VERIFICATION_STILL_VALID

	} else {
		usersContact.VerificationDate = time.Now()
		_, err := pgdb.Model(&usersContact).Where("id = ?", usersContact.Id).Update()
		if err != nil {
			panic(err)
		}
	}
	api := kavenegar.New(lib.SMS_API_KEY)
	sender := lib.SMS_SENDER
	receptor := []string{user}

	message := fmt.Sprintf(`بیرونیستایی عزیز،درخواست شما برای تغییر رمز عبور پذیرفته شد
				. جهت ادامه کد زیر را وارد کنید  
				 %s`, usersContact.Verifier)

	if errorCode < 0 {
		if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
			errorCode = lib.WARNING_SMS_SENT_FAILED
		} else {
			for _, r := range res {
				fmt.Println("MessageID 	= ", r.MessageID)
				fmt.Println("Status    	= ", r.Status)
				errorCode = lib.SMS_SENT_SUCCESSFULLY
			}
		}
	}
	return errorCode
}
