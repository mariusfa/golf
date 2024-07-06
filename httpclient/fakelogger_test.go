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
	UserId        string
}

func (fl *fakeLogger) RequestInfo(requestId string, requestMethod string, requestPath string, requestBody string, userId string) {
	fl.RequestId = requestId
	fl.RequestPath = requestPath
	fl.RequestMethod = requestMethod
	fl.RequestBody = requestBody
	fl.UserId = userId
}

func (fl *fakeLogger) ResponseInfo(requestId string, durationMs string, status int, responseBody string, userId string) {
	fl.RequestId = requestId
	fl.DurationMs = durationMs
	fl.Status = status
	fl.ResponseBody = responseBody
	fl.UserId = userId
}
