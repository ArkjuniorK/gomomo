package schema

type Response struct {
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	Pagination interface{} `json:"pagination"`
}
