package common

type RestResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessRestResult(message interface{}) RestResult {
	return RestResult{
		"200",
		"ok",
		message,
	}
}

func FailedRestResult(code string, message interface{}) RestResult {
	return RestResult{
		code,
		"failed",
		message,
	}
}

func DefaultFailedRestResult(message interface{}) RestResult {
	return RestResult{
		"503",
		"failed",
		message,
	}
}
