package users

import (
	"regexp"
	"time"

	"../../config/db"
	"../../lib"
	"../../model"
	jwt "github.com/dgrijalva/jwt-go"
)

func SanitizeTel(tel string) (string, int) {

	var omitPrefixAreaCode = regexp.MustCompile(`[0]{2}[1-9]{2}|[+][1-9]{2}`)
	var omitStartZero = regexp.MustCompile(`[1-9][0-9]+`)
	var IranianMobileNo = regexp.MustCompile(`[9][0-3][0-9]{8}`)

	tel = omitPrefixAreaCode.ReplaceAllString(tel, "")
	tel = omitStartZero.FindString(tel)
	errorCode := lib.NO_ERROR

	if len(tel) > 10 {
		errorCode = lib.WARNING_TEL_TOO_LONG
	} else if len(tel) < 10 {
		errorCode = lib.WARNING_TEL_TOO_SHORT
	}

	if !IranianMobileNo.MatchString(tel) {
		errorCode = lib.WARNING_INVALID_TEL
	}
	return tel, errorCode
}

func GetUserByContact(contactId string) (model.UsersContacts, int) {
	var usersContacts model.UsersContacts
	errorCode := lib.NO_ERROR
	pgdb := db.DbModel()
	defer pgdb.Close()
	if err := pgdb.Model(&usersContacts).ColumnExpr("users_contacts.*").Join("Inner join users as U on U.id = users_contacts.users_id ").
		Where("users_contacts.address = ?", contactId).Select(); err != nil {
		errorCode = lib.WARNING_CONTACT_NOT_FOUND
	}
	return usersContacts, errorCode

}

func GenerateToken(username string, secret []byte) (string, time.Time) {

	validUntil := time.Now().AddDate(0, 0, 1)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp":  validUntil.Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "not found", validUntil
	}

	return tokenString, validUntil
}
