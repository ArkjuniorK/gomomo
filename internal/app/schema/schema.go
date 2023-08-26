package schema

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Request struct {
	Payload interface{} `json:"payload"`
}
