package users

import (
	"fmt"
	"net/http"
	"os"

	"../../lib"
	"github.com/gin-gonic/gin"

	"github.com/dchest/captcha"
)

func (newUser *NewUser) GetCaptcha() string {
	captchaId := captcha.NewLen(5)
	img, _ := os.Create(fmt.Sprintf("./captcha/%s.png", captchaId))
	captcha.WriteImage(img, captchaId, 200, 70)

	defer img.Close()

	return captchaId
}

func (newUser *NewUser) CaptchaHandler(c *gin.Context) {
	captchaId := newUser.GetCaptcha()
	c.Set("captchaId", captchaId)
	c.JSON(http.StatusOK, gin.H{
		"captcha": captchaId,
	})
}
func (newUser *NewUser) GetAudibleCaptcha(c *gin.Context) {
	errorCode := lib.NO_ERROR
	if err := c.ShouldBind(&newUser); err != nil {
		errorCode = lib.ERROR_CANNOT_GET_POST_PARAMS
	}

	auido, _ := os.Create(fmt.Sprintf("./captcha/%s.wav", newUser.ACaptcha))
	captcha.WriteAudio(auido, newUser.ACaptcha, "en")
	c.JSON(http.StatusOK, gin.H{
		"acaptcha": newUser.ACaptcha,
		"response": lib.ErrorMessages()[errorCode],
		"error":    errorCode,
	})
}
