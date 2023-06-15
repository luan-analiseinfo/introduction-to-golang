package views

type Response struct {
	Code int         `json:"code,omitempty"`
	Body interface{} `json:"body,omitempty"`
}

type PostRequest struct {
	Name string `json:"name,omitempty"`
	Todo string `json:"todo,omitempty"`
}
