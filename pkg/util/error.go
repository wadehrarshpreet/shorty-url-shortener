package util

// errorMap String Constant
//
// key = MODULE_<DESC>
// common key use by multiple packages then drop MODULE_
var errorMap = map[string]string{
	"INVALID_REQUEST_PARAM":   "Invalid Request Params",
	"SOMETHING_WRONG":         "Something Went Wrong! Please try again later.",
	"AUTH_INVALID_USERNAME":   "Username must be between 4 and 32 characters",
	"AUTH_INVALID_EMAIL":      "Invalid Email Address",
	"AUTH_INVALID_PASSWORD":   "Password must be 8 character long with at least one number & one character.",
	"AUTH_USER_ALREADY_EXIST": "Username or Email already taken.",
}
