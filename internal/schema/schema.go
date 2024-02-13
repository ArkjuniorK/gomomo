package schema

// This package contain request and response related schema and not database schema.
// Each package should define their own request schema inside the
// module/package itself.

type Response struct {
	Msg        string      `json:"msg"`
	Code       int         `json:"code"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}
