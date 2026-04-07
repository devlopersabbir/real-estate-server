package networks

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ResCode string

const (
	success_code ResCode = "10000"
	failue_code  ResCode = "10001"
)

type response struct {
	ResCode   ResCode `json:"code" binding:"required"`
	Status    int     `json:"status" binding:"required"`
	Message   string  `json:"message" binding:"required"`
	Timestamp int64   `json:"timestamp"`
	Data      any     `json:"data,omitempty"`
	Errors    any     `json:"errors,omitempty"`
}

func (r *response) GetResCode() ResCode {
	return r.ResCode
}

func (r *response) GetStatus() int {
	return r.Status
}

func (r *response) GetMessage() string {
	return r.Message
}

func (r *response) GetData() any {
	return r.Data
}

type responseSender struct {
	isDebug bool
}

func NewResponseSender(isDebug bool) ResponseSender {
	return &responseSender{isDebug: isDebug}
}

func (s *responseSender) Debug() bool {
	return s.isDebug
}

func (s *responseSender) Send(ctx *gin.Context) SendReponse {
	return &sendResponse{
		ctx:     ctx,
		isDebug: s.isDebug,
	}
}

type sendResponse struct {
	ctx     *gin.Context
	isDebug bool
}

func (s *sendResponse) SuccessMsgResponse(message string) {
	s.ctx.JSON(http.StatusOK, response{
		ResCode:   success_code,
		Status:    http.StatusOK,
		Message:   message,
		Timestamp: time.Now().Unix(),
	})
}

func (s *sendResponse) SuccessDataResponse(message string, data any) {
	s.ctx.JSON(http.StatusOK, response{
		ResCode:   success_code,
		Status:    http.StatusOK,
		Message:   message,
		Timestamp: time.Now().Unix(),
		Data:      data,
	})
}

func (s *sendResponse) handleError(code int, message string, err error) {
	resp := response{
		ResCode:   failue_code,
		Status:    code,
		Message:   message,
		Timestamp: time.Now().Unix(),
	}

	if s.isDebug && err != nil {
		resp.Errors = err.Error()
	}

	s.ctx.JSON(code, resp)
}

func (s *sendResponse) BadRequestError(message string, err error) {
	s.handleError(http.StatusBadRequest, message, err)
}

func (s *sendResponse) ValidationError(message string, errs any) {
	s.ctx.JSON(http.StatusUnprocessableEntity, response{
		ResCode:   failue_code,
		Status:    http.StatusUnprocessableEntity,
		Message:   message,
		Timestamp: time.Now().Unix(),
		Errors:    errs,
	})
}

func (s *sendResponse) ForbiddenError(message string, err error) {
	s.handleError(http.StatusForbidden, message, err)
}

func (s *sendResponse) UnauthorizedError(message string, err error) {
	s.handleError(http.StatusUnauthorized, message, err)
}

func (s *sendResponse) NotFoundError(message string, err error) {
	s.handleError(http.StatusNotFound, message, err)
}

func (s *sendResponse) InternalServerError(message string, err error) {
	s.handleError(http.StatusInternalServerError, message, err)
}

func (s *sendResponse) MixedError(err error) {
	if apiErr, ok := err.(ApiError); ok {
		s.handleError(apiErr.GetCode(), apiErr.GetMessage(), apiErr.Unwrap())
		return
	}
	s.handleError(http.StatusInternalServerError, "Internal Server Error", err)
}

var defaultSender = NewResponseSender(true)

func Send(c *gin.Context) SendReponse {
	return defaultSender.Send(c)
}
