package models

type Result struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"error,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

//swagger Success
type JsonOKResult struct {
	Success bool     `json:"success" example:"true"`
	Message []string `json:"message,omitempty"`
}

//swagger Failure
type JsonFailedResult struct {
	Success bool     `json:"success" example:"false"`
	Error   []string `json:"error,omitempty"`
}

//swagger Success(payload)
type JsonResult struct {
	Success bool        `json:"success" example:"true"`
	Payload interface{} `json:"payload,omitempty" swaggertype:"object"`
}
