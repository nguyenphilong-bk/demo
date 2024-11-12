package utils

type Response struct {
	StatusCode int                    `json:"status_code"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}

type RetrieveResponse struct {
	StatusCode int           `json:"status_code"`
	Message    string        `json:"message,omitempty"`
	Data       []interface{} `json:"data"`
}
