# Httpclient

This is a package to do http requests. Included in these http requests are GET, POST, PUT and DELETE.
All of them have circuit break, bulkhead and transaction logging.

The transaction logging is just an interface. So bring your own logging or use use transaction logging from golf.

## Dependencies
- https://github.com/jarcoal/httpmock (for testing)

## Usage
Get requests
```go
client := NewHttpClient(fakeLogger)
var dto MyDto

requestId := "test" // Used for transaction logging
url := "http://localhost:8080"
headers := map[string]string{"Accept": "application/json"}
getRequest := NewGetRequest(requestId, headers, url)
err := client.GetJson(getRequest, &dto)

```

Post requests
```go
client := NewHttpClient(fakeLogger)
requestDto := MyDto{Name: "Crazy Test"}

requestId := "test"
url := "http://localhost:8080"
headers := map[string]string{"Content-Type": "application/json"}

postRequest := NewPostRequest(requestId, headers, url, requestDto)
err := client.PostJson(postRequest, nil)
```

Put requests
```go
client := NewHttpClient(fakeLogger)
requestDto := MyDto{Name: "Crazy Test"}

requestId := "test"
url := "http://localhost:8080"
headers := map[string]string{"Content-Type": "application/json"}

putRequest := NewPutRequest(requestId, headers, url, requestDto)
err := client.PutJson(putRequest, nil)
```

Delete requests
```go
client := NewHttpClient(fakeLogger)

requestId := "test"
url := "http://localhost:8080/1"
headers := map[string]string{}
getRequest := NewDeleteRequest(requestId, headers, url)
err := client.Delete(getRequest)
```