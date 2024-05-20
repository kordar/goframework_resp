package goframework_resp

import (
	"net/http"
)

func ToResult(c interface{}, t string, message interface{}, data interface{}, count int64) {
	response := container[t]
	if response == nil {
		resultCall(c, http.StatusInternalServerError, Code("fail"), "not found \"response object!\"", nil, 0)
		return
	}

	response.Result(c, message, data, count)
}

func Data(c interface{}, message string, data interface{}, count int64) {
	ToResult(c, SuccessType, message, data, count)
}

func Success(c interface{}, message string, data interface{}) {
	ToResult(c, SuccessType, message, data, -1)
}

func Tenant(c interface{}, message string, data interface{}) {
	ToResult(c, TenantType, message, data, -1)
}

func Error(c interface{}, err interface{}, data interface{}) {
	ToResult(c, ErrorType, err, data, -1)
}

func ErrorV(c interface{}, err interface{}, data interface{}) {
	ToResult(c, ValidErrorType, err, data, -1)
}

func Unauthorized(c interface{}, message string, data interface{}) {
	ToResult(c, UnauthorizedType, message, data, -1)
}

func SuccessOrError(c interface{}, flag bool, successMessage string, err interface{}) {
	if flag {
		ToResult(c, SuccessType, successMessage, nil, -1)
	} else {
		ToResult(c, ErrorType, err, nil, -1)
	}
}

func ErrorOrSuccess(c interface{}, err error) {
	if err == nil {
		ToResult(c, SuccessType, "success", nil, -1)
	} else {
		ToResult(c, ErrorType, err, nil, -1)
	}
}
