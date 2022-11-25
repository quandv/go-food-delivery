package common

type response struct {
	Data   interface{} `json:"data";`
	Paging interface{} `json:"paging,omitempty";`
	Filter interface{} `json:"filter,omitempty";`
}

func NewResponse(data, paging, filter interface{}) *response {
	return &response{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}
