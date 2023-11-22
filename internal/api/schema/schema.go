package schema

type Request struct {
	Payload interface{}
}

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
