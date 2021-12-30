package common

import (
	"fmt"
	"net/http"
)

type ErrorID string

type AppError struct {
	errorID    ErrorID                // Unique error ID
	message    string                 // Customizable error message
	statusCode int                    // Fixed HTTP status code for each ErrorID
	data       map[string]interface{} // Error data, mostly for debugging purpose
}

func (e *AppError) ErrorID() ErrorID {
	return e.errorID
}

func (e *AppError) Error() string {
	return fmt.Sprintf("err: %s, statusCode: %d, message: %s, data: %v", e.errorID, e.statusCode, e.message, e.data)
}

func (e *AppError) StatusCode() int {
	return e.statusCode
}

func (e *AppError) With(message string, data map[string]interface{}) *AppError {
	return &AppError{
		errorID:    e.errorID,
		message:    message,
		statusCode: e.statusCode,
		data:       data,
	}
}

func (e *AppError) WithMessage(message string) *AppError {
	return &AppError{
		errorID:    e.errorID,
		message:    message,
		statusCode: e.statusCode,
		data:       e.data,
	}
}

func (e *AppError) WithData(data map[string]interface{}) *AppError {
	return &AppError{
		errorID:    e.errorID,
		message:    e.message,
		statusCode: e.statusCode,
		data:       data,
	}
}

func (e *AppError) Is(err error) bool {
	appErr, ok := err.(*AppError)
	if !ok {
		return false
	}
	return appErr.ErrorID() == e.ErrorID()
}

func newAppError(errID ErrorID, message string, statusCode int) *AppError {
	return &AppError{
		errorID:    errID,
		message:    message,
		statusCode: statusCode,
		data:       map[string]interface{}{},
	}
}

var (
	cgkCoinNotFoundID ErrorID = "cgk_coin_not_found"
	CgkCoinNotFound           = newAppError(cgkCoinNotFoundID, "coin id not found", http.StatusNotFound)

	cgkInvalidDataID ErrorID = "cgk_invalid_data"
	CgkInvalidData           = newAppError(cgkInvalidDataID, "invlaid data", http.StatusNotFound)

	cgkRateLimitID ErrorID = "cgk_rate_limit"
	CgkRateLimit           = newAppError(cgkRateLimitID, "hit rate limit", http.StatusTooManyRequests)

	nodeErrorID ErrorID = "node_error"
	NodeError           = newAppError(nodeErrorID, "node error", http.StatusInternalServerError)
)
