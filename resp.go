package goframework_resp

import (
	"net/http"
)

var (
	SuccessType      = "success"
	ErrorType        = "error"
	ValidErrorType   = "valid-error"
	OutputType       = "output"
	UnauthorizedType = "unauthorized"
	TenantType       = "tenant"
)

var container = map[string]IResponse{
	"success": SuccessResult{},
	"error":   ErrorResult{},
}

func RegRespFunc(name string, response IResponse) {
	container[name] = response
}

type IResponse interface {
	HttpStatus() int
	Result(c interface{}, message interface{}, data interface{}, count int64)
}

// SuccessResult 成功
type SuccessResult struct {
}

func (s SuccessResult) HttpStatus() int {
	return http.StatusOK
}

func (s SuccessResult) Result(c interface{}, message interface{}, data interface{}, count int64) {
	if value, ok := message.(string); ok && value != "" {
		resultCall(c, s.HttpStatus(), Code("success"), value, data, count)
		return
	}

	resultCall(c, s.HttpStatus(), Code("success"), "success", data, count)
}

// ----------- Error ------------

type ErrorResult struct {
}

func (e ErrorResult) HttpStatus() int {
	return http.StatusOK
}

func (e ErrorResult) Result(c interface{}, message interface{}, data interface{}, count int64) {
	//TODO implement me
	if value, ok := message.(string); ok && value != "" {
		resultCall(c, e.HttpStatus(), Code("error"), value, data, count)
		return
	}

	if value, ok := message.(error); ok {
		resultCall(c, e.HttpStatus(), Code("error"), value.Error(), data, count)
		return
	}

	resultCall(c, e.HttpStatus(), Code("error"), "error", data, count)

}

// -------------- Output Excel -------------------------

type OutputResponse struct {
}

func (o OutputResponse) HttpStatus() int {
	return http.StatusOK
}

func (o OutputResponse) Result(c interface{}, message interface{}, data interface{}, count int64) {
	// TODO implement me

	/*
		if value, ok := data.(IOutput); ok && value != nil {
			if value.Type() == "browser" {
				// output web
				for k, v := range value.Header() {
					c.Header(k, v)
				}
				_, _ = c.Writer.Write(value.Data())
				return
			}

			msg := "output success!"
			if tmessage, found := message.(string); found && tmessage != "" {
				msg = tmessage
			}

			c.JSON(http.StatusOK, map[string]interface{}{
				"code": Code("success"), "message": msg, "data": value.Data(), "params": value.Params(),
			})
			return
		}
	*/

	resultCall(c, o.HttpStatus(), Code("success"), "output success!", nil, -1)
}

// -------------- Unauthorized Excel -------------------------

type UnauthorizedResponse struct {
}

func (o UnauthorizedResponse) HttpStatus() int {
	return http.StatusOK
}

func (o UnauthorizedResponse) Result(c interface{}, message interface{}, data interface{}, count int64) {
	resultCall(c, o.HttpStatus(), Code("unauthorized"), "unauthorized", nil, -1)
}

// -------------- MultiTenant -------------------------

type MultiTenantResponse struct {
}

func (o MultiTenantResponse) HttpStatus() int {
	return http.StatusOK
}

func (o MultiTenantResponse) Result(c interface{}, message interface{}, data interface{}, count int64) {
	resultCall(c, o.HttpStatus(), Code("multitenant"), "multiTenant", data, -1)
}
