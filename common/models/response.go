package models

type Response struct {
	// code
	Code int `json:"code" example:"200"`
	// data
	Data interface{} `json:"data"`
	// msg
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

// ReturnOK
func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

// ReturnError
func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}
