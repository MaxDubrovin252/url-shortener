package models

type Request struct {
	URL   string `json:"url" binding:"required"`
	Alias string `json:"alias" binding:"required"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "error"
)

func OK() *Response {
	return &Response{
		Status: StatusOK,
	}

}

func Error(message string) *Response {
	return &Response{
		Status: StatusError,
		Error:  message,
	}
}
