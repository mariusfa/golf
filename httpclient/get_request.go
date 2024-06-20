package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type GetRequest struct {
	RequestId string
	Headers   map[string]string
	Url       string
}

func NewGetRequest(requestId string, headers map[string]string, url string) *GetRequest {
	return &GetRequest{RequestId: requestId, Headers: headers, Url: url}
}

func (c *HttpClient) GetJson(request *GetRequest, dto any) error {
	httpAction := func() error {
		return c.getJsonPlain(request, dto)
	}
	bulkheadAction := func() error {
		return c.bulkhead.Execute(httpAction)
	}
	circuitbreakerAction := func() error {
		return c.circuitbreaker.Execute(bulkheadAction)
	}
	return circuitbreakerAction()
}

func (c *HttpClient) getJsonPlain(request *GetRequest, dto any) error {
	req, err := http.NewRequest("GET", request.Url, nil)
	if err != nil {
		return err
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	c.logger.RequestInfo(request.RequestId, "GET", request.Url, "")
	start := time.Now()
	resp, err := c.client.Do(req)
	duratonMs := time.Since(start).Milliseconds()
	if err != nil {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), 0, "")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "")
		return errors.New(fmt.Sprintf("%s from %s", resp.Status, request.Url))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "")
		return err
	}

	bodyString := string(bodyBytes)

	c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, bodyString)

	bodyReader := bytes.NewReader(bodyBytes)

	return json.NewDecoder(bodyReader).Decode(dto)
}
