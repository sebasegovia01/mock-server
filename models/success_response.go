package models

type SuccessResponse struct {
	Result struct {
		Status      ResultStatus `json:"status"`
		Description string       `json:"description,omitempty"`
		Data        interface{}  `json:"data,omitempty"`
	} `json:"Result"`
}
