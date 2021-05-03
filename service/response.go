package service

// Response API response structure
type Response struct {
	Code   string `json:"code,omitempty" example:"internalError"`
	Status string `json:"status,omitempty" example:"success"`
}

// New returns service error code
func (r Response) error(code string) Response {
	r.Code = code
	return r
}

func (r Response) internal() Response {
	r.Code = "internalError"
	return r
}

func (r Response) success() Response {
	r.Status = "success"
	return r
}
