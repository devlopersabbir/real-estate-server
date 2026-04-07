package networks

import "github.com/gin-gonic/gin"

type ApiError interface {
	GetCode() int
	GetMessage() string
	Error() string
	Unwrap() error
}

type Response interface {
	GetResCode() ResCode
	GetStatus() int
	GetMessage() string
	GetData() any
}

type SendReponse interface {
	SuccessMsgResponse(message string)
	SuccessDataResponse(message string, data any)
	BadRequestError(message string, err error)
	ValidationError(message string, errs any)
	ForbiddenError(message string, err error)
	UnauthorizedError(message string, err error)
	NotFoundError(message string, err error)
	InternalServerError(message string, err error)
	MixedError(err error)
}

type ResponseSender interface {
	Debug() bool
	Send(ctx *gin.Context) SendReponse
}
