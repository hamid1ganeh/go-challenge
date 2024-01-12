package response

type Response struct {
	Success bool
	Message string
	Data    interface{}
}

func Success(Data interface{}, message string) Response {
	if message == "" {
		message = "Success"
	}
	return Response{
		Success: true,
		Message: message,
		Data:    Data,
	}
}

func Fail(Data interface{}, message string) Response {
	if message == "" {
		message = "Fail"
	}
	return Response{
		Success: true,
		Message: message,
		Data:    Data,
	}
}
func NotFound(Data interface{}, message string) Response {
	if message == "" {
		message = "NotFound"
	}
	return Response{
		Success: false,
		Message: message,
		Data:    Data,
	}
}

func Conflict(Data interface{}, message string) Response {
	if message == "" {
		message = "Dose exist"
	}
	return Response{
		Success: false,
		Message: message,
		Data:    Data,
	}
}
