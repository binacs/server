package base

import (
	"regexp"
)

var (
	//regNumber, _     = regexp.Compile("^[0-9]+$")
	regPhone, _      = regexp.Compile(`^1[\d]{10}$`)
	regEmail, _      = regexp.Compile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	regVerifyCode, _ = regexp.Compile(`^[a-zA-Z0-9]{6}$`)
	regPasswd, _     = regexp.Compile(`^[0-9a-z]{32}$`)

	//regHexWith0x, _     = regexp.Compile("^[0][x][0-9a-f]+$")
	//regHexWithout0x, _  = regexp.Compile("^[0-9a-f]+$")
	//regAddressWith0x, _ = regexp.Compile("^[0][x][0-9a-f]{40}$")
	regBase64, _ = regexp.Compile("^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{4}|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)$")
)

// IsPhone phone
func IsPhone(phone string) bool {
	return regPhone.MatchString(phone)
}

// IsEmail email
func IsEmail(email string) bool {
	return regEmail.MatchString(email)
}

// IsVerifyCode verify-code
func IsVerifyCode(verifyCode string) bool {
	return regVerifyCode.MatchString(verifyCode)
}

// IsPassword password
func IsPassword(passwod string) bool {
	return regPasswd.MatchString(passwod)
}

// IsBase64 base64
func IsBase64(b64 string) bool {
	return regBase64.MatchString(b64)
}
