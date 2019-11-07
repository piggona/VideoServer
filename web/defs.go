package main

type ApiBody struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	ReqBody string `json:"req_body"`
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"errorCode"`
}

var (
	ErrorRequestNotRecognized = Err{Error: "api not recognized, bad request", ErrorCode: "001"}
	ErrorRequestParseFailed   = Err{Error: "request body is not correct", ErrorCode: "002"}
	ErrorInternalFaults       = Err{Error: "internal error", ErrorCode: "003"}
)
