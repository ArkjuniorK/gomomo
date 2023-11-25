package schema

// This package contain request related schema and not database schema.
// Each package should define their own schema in separate files.

type Request struct {
	Payload interface{}
}

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
