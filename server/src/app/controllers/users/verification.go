package users

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"../../config/db"
	"../../lib"
	"../../model"
	"github.com/gin-gonic/gin"
	kavenegar "github.com/kavenegar/kavenegar-go"
)

func (mw *AuthMiddleware) SendVerificationCodeHandler(c *gin.Context) {

	contact, errorCode := mw.UpdateVerificationCode(mw.Claim["user"].(string))
	api := kavenegar.New(lib.SMS_API_KEY)
	sender := "10004346"
	receptor := []string{mw.Claim["user"].(string)}

	message := fmt.Sprintf(`بیرونیستایی عزیز، این پیام حاوی کد فعال سازی شما 
				 است. لحظات خوشی را برای شما آرزومندیم  
				 %s`, contact.Verifier)

	if errorCode < 0 {
		if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
			switch err := err.(type) {
			case *kavenegar.APIError:
				fmt.Println(err.Error())
				errorCode = lib.WARNING_SMS_SENT_FAILED
			case *kavenegar.HTTPError:
				fmt.Println(err.Error())
				errorCode = lib.WARNING_SMS_SENT_FAILED
			default:
				fmt.Println(err.Error())
				errorCode = lib.WARNING_SMS_SENT_FAILED
			}
		} else {
			for _, r := range res {
				fmt.Println("MessageID 	= ", r.MessageID)
				fmt.Println("Status    	= ", r.Status)
				errorCode = lib.SMS_SENT_SUCCESSFULLY
				//...
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"credentials": mw.Claim,
		"error":       errorCode,
		"response":    lib.ErrorMessages()[errorCode],
	})

}

func (mw *AuthMiddleware) UpdateVerificationCode(contactId string) (model.UsersContacts, int) {
	errorCode := lib.NO_ERROR
	usersContacts, _ := GetUserByContact(contactId)

	pgdb := db.DbModel()
	defer pgdb.Close()
	usersContacts.Verifier = fmt.Sprintf("%d", 100000+rand.Intn(899999))

	diff := time.Now().Sub(usersContacts.VerificationDate)

	if int(diff.Seconds()) < lib.VERIFICATION_EXPIRATION_SECONDS {
		errorCode = lib.WARNING_LAST_VERIFICATION_STILL_VALID
	} else {
		usersContacts.VerificationDate = time.Now()
		_, err := pgdb.Model(&usersContacts).Where("id = ?", usersContacts.Id).Update()
		if err != nil {
			panic(err)
		}
	}

	return usersContacts, errorCode
}

func (mw *AuthMiddleware) ContactVerificationByCodeHandler(c *gin.Context) {
	var verify VerifyUser
	errorCode := lib.NO_ERROR
	if err := c.ShouldBind(&verify); err != nil {
		panic("Cannot bind")
	}
	usersContacts, _ := GetUserByContact(mw.Claim["user"].(string))
	if usersContacts.Verifier == verify.Verification {
		errorCode = lib.VERIFICATION_SUCCESSFUL
		pgdb := db.DbModel()
		defer pgdb.Close()
		usersContacts.IsVerified = true
		pgdb.Model(&usersContacts).Where("id = ?", usersContacts.Id).Update()
	} else {
		errorCode = lib.WARNING_VERIFICATION_FAILED
	}

	c.JSON(http.StatusOK, gin.H{
		"credentials": mw.Claim,
		"error":       errorCode,
		"response":    lib.ErrorMessages()[errorCode],
	})

}
