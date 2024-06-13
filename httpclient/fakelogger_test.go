package httpclient

type MyDto struct {
	Name string `json:"name"`
}

type fakeLogger struct {
	RequestId     string
	RequestPath   string
	RequestMethod string
	DurationMs    string
	Status        int
	ResponseBody  string
	RequestBody   string
}

func (fl *fakeLogger) RequestInfo(requestId string, requestMethod string, requestPath string, requestBody string) {
	fl.RequestId = requestId
	fl.RequestPath = requestPath
	fl.RequestMethod = requestMethod
	fl.RequestBody = requestBody
}

func (fl *fakeLogger) ResponseInfo(requestId string, durationMs string, status int, responseBody string) {
	fl.RequestId = requestId
	fl.DurationMs = durationMs
	fl.Status = status
	fl.ResponseBody = responseBody
}
