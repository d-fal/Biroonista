package users

import (
	"net/http"
	"strings"
	"time"

	"../../config/db"
	"../../lib"
	"../../model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	Token  string
	Bearer string
	Claim  jwt.MapClaims
}

type ClaimedUser struct {
	UserName  string
	FirstName string
	LastName  string
	Password  string
}

func AuthUser(userInstance ClaimedUser) (bool, bool) {

	pgdb := db.DbModel()
	defer pgdb.Close()

	var usersContacts model.UsersContacts

	err := pgdb.Model(&usersContacts).
		Join("Inner Join users as U on U.id = users_contacts.users_id").ColumnExpr("users_contacts.*").
		Where("users_contacts.address = ? and U.password = crypt(?, password)",
			userInstance.UserName, userInstance.Password).Select()

	if err != nil {
		return false, false
	}
	return true, usersContacts.IsVerified

}
func (mw *AuthMiddleware) Unauthorized(claim jwt.MapClaims, c *gin.Context) {

	var errorCode int
	if expiry, ok := claim["exp"]; ok {
		if int64(expiry.(float64)) < time.Now().Unix() {
			errorCode = lib.ERROR_TOKEN_EXPIRED
		}
	} else {
		errorCode = lib.ERROR_TOKEN_NOT_VALID
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error":    errorCode,
		"response": lib.ErrorMessages()[errorCode],
	})
	c.AbortWithStatus(http.StatusUnauthorized)
}

func (mw *AuthMiddleware) RefreshTokenHandler(c *gin.Context) {
	if isAuth, claim := mw.Authenticate(c); isAuth {
		var user ClaimedUser
		user.UserName = claim["user"].(string)
		token, exp := GenerateToken(user.UserName, lib.Secret)
		mw.Token = token
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"exp":   exp.Unix(),
		})
	}

}

func (mw *AuthMiddleware) Authenticate(c *gin.Context) (bool, jwt.MapClaims) {
	bearer := c.GetHeader("Authorization")
	mw.Bearer = bearer
	var token string
	slices := strings.Split(bearer, " ")

	if slices[0] == "Bearer" {
		if len(slices) == 2 {
			token = slices[1]
		}
	}
	return mw.VerifyToken(token)
}
func (mw *AuthMiddleware) AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAuth, claim := mw.Authenticate(c)
		if !isAuth {
			mw.Unauthorized(claim, c)
		} else {
			c.Set("user", claim)
			mw.Claim = claim

		}
	}
}

func (mw *AuthMiddleware) VerifyToken(tokenString string) (bool, jwt.MapClaims) {
	if tokenString == "" {
		return false, nil
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, nil
		}
		return lib.Secret, nil
	})

	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mw.Token = tokenString
		return true, claim
	} else {
		return false, claim
	}

}
func (mw *AuthMiddleware) LogoutHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"error":    lib.LOGOUT_SUCCESSFUL,
			"response": lib.ErrorMessages()[lib.LOGOUT_SUCCESSFUL],
		},
	)
}

func (mw *AuthMiddleware) MarshalBasicInfoHandler(c *gin.Context) {
	userContact, _ := GetUserByContact(mw.Claim["user"].(string))
	errorCode := lib.NO_ERROR

	if !userContact.IsVerified {
		errorCode = lib.WARNING_CONTACT_NOT_VERIFIED
	}
	c.JSON(http.StatusOK, gin.H{
		"credentials": mw.Claim,
		"error":       errorCode,
		"response":    lib.ErrorMessages()[errorCode],
	})
}
