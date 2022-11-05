package transfer

import constants "github.com/devnura/go-echo-rest-api/constants"

// base transfer
type Base struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// error transfer
type CustomError struct {
	Base
}

// success transfer
type Success struct {
	Base
	Data interface{} `json:"data"`
}

// method of error transfer
func (e *CustomError) Error() string {
	return e.Code + ": " + e.Message
}

func NewCustomError(code string) *CustomError {
	return &CustomError{
		Base{
			Code: code,
		},
	}
}

func NewCustomErrorMsg(code string, msg string) *CustomError {
	return &CustomError{
		Base{
			Code:    code,
			Message: msg,
		},
	}
}

func NewSuccess(data interface{}) *Success {
	sucess := Success{}
	sucess.Code = constants.SUCCESS
	sucess.Data = data
	return &sucess
}
