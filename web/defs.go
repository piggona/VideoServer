package main

type ApiBody struct {
	Url string `json:"url"`
	// Url, 此处的Url应当是一个服务器的代号或子域名，
	// 然后也许在执行request函数的过程中需要再执行判断一次,在后端得到最终访问的明确url
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
