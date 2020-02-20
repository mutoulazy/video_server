package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpStatusCode int
	Error          Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpStatusCode: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrorResponse{HttpStatusCode: 401, Error: Err{Error: "User authentication failed", ErrorCode: "002"}}
	ErrorDBError                = ErrorResponse{HttpStatusCode: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFaults         = ErrorResponse{HttpStatusCode: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
