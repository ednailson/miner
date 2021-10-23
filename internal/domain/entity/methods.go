package entity

const (
	JsonrpcVersion  = "2.0"
	AuthorizeMethod = "mining.authorize"
	SubscribeMethod = "mining.subscribe"
)

type Success struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      uint64      `json:"id"`
}

type Fail struct {
	Jsonrpc string  `json:"jsonrpc"`
	Error   Error   `json:"error"`
	ID      *uint64 `json:"id"`
}

func NewSuccess(id uint64, result interface{}) Success {
	return Success{
		Jsonrpc: JsonrpcVersion,
		Result:  result,
		ID:      id,
	}
}

func NewFail(id *uint64, e Error) Fail {
	return Fail{
		Jsonrpc: JsonrpcVersion,
		Error:   e,
		ID:      id,
	}
}
