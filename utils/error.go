package utils

type HttpError interface {
    Error() string
    StatusCode() int
    UserMessage() string
}

type ClientError struct {
    Code    int
    Message string
    UserMsg string
}

func (e *ClientError) Error() string {
    return e.Message
}

func (e *ClientError) StatusCode() int {
    return e.Code
}

func (e *ClientError) UserMessage() string {
    return e.UserMsg
}

func NewClientError(code int, message, userMessage string) *ClientError {
    return &ClientError{
        Code:    code,
        Message: message,
        UserMsg: userMessage,
    }
}

type ServerError struct {
    Code    int
    Message string
    UserMsg string
}

func (e *ServerError) Error() string {
    return e.Message
}

func (e *ServerError) StatusCode() int {
    return e.Code
}

func (e *ServerError) UserMessage() string {
    return e.UserMsg
}

func NewServerError(code int, message, userMessage string) *ServerError {
    return &ServerError{
        Code:    code,
        Message: message,
        UserMsg: userMessage,
    }
}