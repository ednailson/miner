package entity

type Error struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
}

func ErrorMethodNotFound() Error {
	return Error{
		Code:    -32601,
		Message: "Method not found",
	}
}

func ErrorInvalidParams() Error {
	return Error{
		Code:    -32602,
		Message: "Invalid params",
	}
}

func ErrorInvalidRequest() Error {
	return Error{
		Code:    -32600,
		Message: "Invalid Request",
	}
}

func ErrorServer() Error {
	return Error{
		Code:    -32000,
		Message: "Server error",
	}
}

func ErrorParse() Error {
	return Error{
		Code:    -32700,
		Message: "Parse error",
	}
}
