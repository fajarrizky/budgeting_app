package exception

import "net/http"

type HttpException interface {
	Exception
	Status() int
}

type httpException struct {
	Exception
	status int
}

func (ex *httpException) Status() int {
	return ex.status
}

func NewHttpException(status int, ex Exception) HttpException {

	return &httpException{
		Exception: ex,
		status:    status,
	}

}

func ToHttpExceptionFromError(err error) HttpException {

	if errVal, ok := err.(HttpException); ok {
		return errVal
	}

	return InternalServerException(err)
}

func InternalServerException(err error, msg ...string) HttpException {

	ex := NewHttpException(
		http.StatusInternalServerError,
		New(
			INTERNAL_SERVER_ERROR,
			getString(firstOne(msg), "Internal Server Error"), err),
	)

	return ex
}

func BadRequestException(err error, msg ...string) HttpException {

	ex := NewHttpException(
		http.StatusBadRequest,
		New(
			BAD_REQUEST_ERROR,
			getString(firstOne(msg), "Bad request"), err),
	)

	return ex
}

func NotFoundException(err error, msg ...string) HttpException {

	ex := NewHttpException(
		http.StatusNotFound,
		New(
			NOT_FOUND_ERROR,
			getString(firstOne(msg), "resource not found"), err),
	)

	return ex
}

func NotAllowedException(err error, msg ...string) HttpException {

	ex := NewHttpException(
		http.StatusMethodNotAllowed,
		New(
			NOT_ALLOWED_ERROR,
			getString(firstOne(msg), "not allowed"), err),
	)

	return ex
}

func ForbiddenException(err error, msg ...string) HttpException {

	ex := NewHttpException(
		http.StatusForbidden,
		New(
			FORBIDDEN_ERROR,
			getString(firstOne(msg), "forbidden access"), err),
	)

	return ex
}
