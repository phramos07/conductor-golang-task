package model

// CustomError Interface
type CustomError interface {
	error
	StatusCode() int
}

type requestError struct {
	CustomError

	errStr     string
	statusCode int
}

// NewrequestError Makes new custom request with given status code and str
func NewrequestError(errStr string, statusCode int) CustomError {
	return &requestError{
		statusCode: statusCode,
		errStr:     errStr,
	}
}

func (r *requestError) Error() string {
	return r.errStr
}

func (r *requestError) StatusCode() int {
	return r.statusCode
}
