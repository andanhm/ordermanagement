package errors

import "fmt"

type (
	// Code is the error code enum.
	Code string

	// Error is the interface providing the implementation needed to return error
	Error interface {
		// Code is the error code and must have one of the values defines by the package
		// constants.
		Code() Code

		// Description is a human-readable ASCII text providing additional information, used
		// to assist the client developer in understanding the error that occurred.
		Description() string

		// Error returns the underlying error's message.
		Error() string
	}

	// anglebrokerError is a simple implementation of Error.
	anglebrokerError struct {
		ErrorCode        Code   `json:"error"`
		ErrorDescription string `json:"error_description,omitempty"`
	}
)

// New creates an error suitable to be returned in the body of OAuth2 error responses.
func New(code Code, description string) Error {
	return &anglebrokerError{code, description}
}

// oAuth2Error implements Error.
func (e *anglebrokerError) Code() Code { return e.ErrorCode }
func (e *anglebrokerError) Description() string {
	return fmt.Sprintf("%v: %s", e.ErrorCode, e.ErrorDescription)
}
func (e *anglebrokerError) Error() string { return e.ErrorDescription }
