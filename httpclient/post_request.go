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

type PostRequest struct {
	RequestId string
	Headers   map[string]string
	Url       string
	Body      any
}

func NewPostRequest(requestId string, headers map[string]string, url string, body any) *PostRequest {
	return &PostRequest{RequestId: requestId, Headers: headers, Url: url, Body: body}
}

func (c *HttpClient) PostJson(request *PostRequest, responseDto any) error {
	httpAction := func() error {
		return c.postJsonPlain(request, responseDto)
	}
	bulkheadAction := func() error {
		return c.bulkhead.Execute(httpAction)
	}
	circuitbreakerAction := func() error {
		return c.circuitbreaker.Execute(bulkheadAction)
	}
	return circuitbreakerAction()
}

func (c *HttpClient) postJsonPlain(request *PostRequest, responseDto any) error {
	body, err := json.Marshal(request.Body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", request.Url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	c.logger.RequestInfo(request.RequestId, "POST", request.Url, string(body))
	start := time.Now()
	resp, err := c.client.Do(req)
	duratonMs := time.Since(start).Milliseconds()
	if err != nil {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), 0, "")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "")
		return errors.New(fmt.Sprintf("%s from %s", resp.Status, request.Url))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "")
		return err
	}

	if len(bodyBytes) == 0 {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "")
		return nil
	}

	err = json.Unmarshal(bodyBytes, responseDto)
	if err != nil {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, string(bodyBytes))
		return err
	}

	c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, string(bodyBytes))
	return nil
}
