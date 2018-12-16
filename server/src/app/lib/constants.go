package lib

import (
	"fmt"
)

var (
	NO_ERROR          = -1
	LOGOUT_SUCCESSFUL = 1

	LOGIN_SUCCESSFUL = 3

	ERROR_TOKEN_NOT_VALID = 506
	ERROR_TOKEN_EXPIRED   = 507
	ERROR_NO_TOKEN        = 508
	ERROR_LOGOUT_FAILED   = 502
	ERROR_BINDING_FAILED  = 503

	WARNING_PASSWORDS_DOES_NOT_MATCH           = 109
	WARNING_PASSWORD_TOO_SHORT                 = 100
	WARNING_PASSWORD_TOO_WEAK                  = 119
	WARNING_USER_ALREADY_EXISTS                = 102
	WARNING_CANNOT_SAVE_CONTACT                = 103
	WARNING_INVALID_TEL                        = 104
	WARNING_TEL_TOO_SHORT                      = 105
	WARNING_TEL_TOO_LONG                       = 106
	WARNING_INCORRECT_CAPTCHA                  = 107
	ERROR_CANNOT_GET_POST_PARAMS               = 108
	WARNING_CONTACT_NOT_VERIFIED               = 101
	WARNING_LOGIN_FAILED_INCORRECT_CREDENTIALS = 110
	WARNING_LOGIN_FAILED_USER_NOT_FOUND        = 111
	WARNING_VERIFICATION_FAILED                = 112
	WARNING_LAST_VERIFICATION_STILL_VALID      = 113
	WARNING_SMS_SENT_FAILED                    = 114
	WARNING_CONTACT_NOT_FOUND                  = 115
	WARNING_SUSCPICIOUS_USER                   = 116
	WARNING_INVALID_VERIFICATION_PATTERN       = 117
	WARNING_CANNOT_CHANGE_PASSWORD             = 118

	VERIFICATION_SUCCESSFUL = -901
	RESET_TOKEN_GENERATED   = -902
	SMS_SENT_SUCCESSFULLY   = -903
	PATTERN_ACCEPTED        = -904
	PASSWORD_CHANGED        = -905

	TYPE_CONTACT_MOBILE = 1
)
var errorMessage = map[int]string{
	NO_ERROR: "Success",
	WARNING_LOGIN_FAILED_INCORRECT_CREDENTIALS: "Incorrect tel/email or password",
	WARNING_LOGIN_FAILED_USER_NOT_FOUND:        "User not found",
	WARNING_PASSWORDS_DOES_NOT_MATCH:           "Passwords do not match",
	WARNING_USER_ALREADY_EXISTS:                "User already exists",
	WARNING_PASSWORD_TOO_WEAK:                  "Password too weak",
	WARNING_PASSWORD_TOO_SHORT:                 "Password too short",
	WARNING_INVALID_TEL:                        "The provided number is not valid",
	WARNING_TEL_TOO_SHORT:                      "Provided mobile number is shorter than it should be",
	WARNING_TEL_TOO_LONG:                       "Number is longer than standard format",
	WARNING_INCORRECT_CAPTCHA:                  "Incorrect captcha",
	WARNING_CANNOT_SAVE_CONTACT:                "Contact cannot be saved",
	ERROR_CANNOT_GET_POST_PARAMS:               "Cannot load post items",
	LOGOUT_SUCCESSFUL:                          "Log-out was successfuk",
	ERROR_NO_TOKEN:                             "Token is not provided by the client",
	ERROR_TOKEN_EXPIRED:                        "Token is not valid",
	ERROR_TOKEN_NOT_VALID:                      "Invalid token",
	WARNING_CONTACT_NOT_VERIFIED:               "Number is not verified",
	VERIFICATION_SUCCESSFUL:                    "Verification accomplished",
	WARNING_VERIFICATION_FAILED:                "Verification failed, incorrect code provided",
	WARNING_LAST_VERIFICATION_STILL_VALID: fmt.Sprintf(`e have already sent a message to you, please try again after %d seconds`,
		VERIFICATION_EXPIRATION_SECONDS),
	RESET_TOKEN_GENERATED:                "Token generated",
	ERROR_BINDING_FAILED:                 "Binding failed",
	SMS_SENT_SUCCESSFULLY:                "SMS sent successfully",
	WARNING_SMS_SENT_FAILED:              "SMS was not sent",
	WARNING_CONTACT_NOT_FOUND:            "Contact was not found",
	WARNING_SUSCPICIOUS_USER:             "Sucpicious action by user",
	PATTERN_ACCEPTED:                     "Verification code accepted",
	WARNING_INVALID_VERIFICATION_PATTERN: "Invalid verification pattern",
	WARNING_CANNOT_CHANGE_PASSWORD:       "Cannot change password",
	PASSWORD_CHANGED:                     "Password changed",
}

func ErrorMessages() map[int]string {

	return errorMessage
}
