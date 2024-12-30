package types

type ApiResponse struct {
	Result      int64       `json:"resultCode"`
	Description string      `json:"description"`
	ErrCode     interface{} `json:"errCode"`
}

func NewApiResponse(description string, result int64, errCode interface{}) *ApiResponse {
	return &ApiResponse{
		Result:      result,
		Description: description,
		ErrCode:     errCode,
	}
}
