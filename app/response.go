package app

type Response struct {
	Msg  string      `json:"message"`
	Code int         `json:"code,string"`
	Data interface{} `json:"data,omitempty"`
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Msg:  "success",
		Code: 0,
		Data: data,
	}
}

func ErrorResponse(err Err) Response {
	return Response{
		Msg:  err.Message,
		Code: err.Code,
		Data: nil,
	}
}
