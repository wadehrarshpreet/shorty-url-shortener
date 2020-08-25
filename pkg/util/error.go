package util

// errorMap String Constant
//
// key = MODULE_<DESC>
// common key use by multiple packages then drop MODULE_

type errorData struct {
	errCode int
	message string
}

var errorMap = map[string]errorData{
	"INVALID_REQUEST_PARAM":    {10001, "Invalid Request Params"},
	"SOMETHING_WRONG":          {10002, "Something Went Wrong! Please try again later"},
	"AUTH_INVALID_USERNAME":    {10003, "Username must be between 4 and 32 characters"},
	"AUTH_INVALID_EMAIL":       {10004, "Invalid Email Address"},
	"AUTH_INVALID_PASSWORD":    {10005, "Password must be 8 character long with at least one number & one character"},
	"AUTH_USER_ALREADY_EXIST":  {10006, "Username already taken"},
	"AUTH_EMAIL_ALREADY_EXIST": {10007, "Email Already exist"},
	"AUTH_FAILED":              {10009, "Invalid Username or Password"},
}
