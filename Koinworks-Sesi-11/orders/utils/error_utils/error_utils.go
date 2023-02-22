package error_utils

import "net/http"

type MessageErr interface {
	Message() string
	Status() int
	Error() string
}

type MessageErrData struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e *MessageErrData) Message() string {
	return e.ErrMessage
}

func (e *MessageErrData) Status() int {
	return e.ErrStatus
}

func (e *MessageErrData) Error() string {
	return e.ErrError
}

func NewNotFoundError(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func NewNotAuthenticated(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "not_authenticated",
	}
}

func NewNotAuthorized(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusForbidden,
		ErrError:   "not_authorized",
	}
}

func NewBadRequest(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewInternalServerErrorr(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "server_error",
	}
}

func NewUnprocessibleEntityError(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "invalid_request",
	}
}
