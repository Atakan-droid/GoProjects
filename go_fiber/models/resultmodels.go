package models

type ApiResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
}

func CreateSuccess(data interface{}) *ApiResult {
	return &ApiResult{
		Success: true,
		Data:    data,
	}
}

func CreateError(errors []string) *ApiResult {
	return &ApiResult{
		Success: false,
		Errors:  errors,
	}
}
