package errors

type HTTPError struct {
	Code    int
	Message string
}

// NewHTTPError returns a new HTTPError with the given code and message.
func NewHTTPError(code int, message string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: message,
	}
}

// Error returns the error message.
func (e HTTPError) Error() string {
	return e.Message
}
