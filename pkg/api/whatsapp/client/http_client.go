package whatsapp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	loggerUtils "github.com/jkandasa/whatsapp-cloud-api/pkg/utils/logger"

	"go.uber.org/zap"
)

const (
	requestContentType = "application/json"
)

type Client struct {
	logger     *zap.Logger
	baseURL    string
	headers    map[string]string
	httpClient *http.Client
}

func New(ctx context.Context, baseURL string, headers map[string]string) (*Client, error) {
	logger, err := loggerUtils.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	_client := Client{
		logger:     logger.Named("custom_client"),
		baseURL:    baseURL,
		headers:    headers,
		httpClient: &http.Client{},
	}
	return &_client, nil
}

// converts the given interface to map with json tag
func toMap(data any) (map[string]any, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := map[string]any{}
	err = json.Unmarshal(bytes, &mapData)
	if err != nil {
		return nil, err
	}

	return mapData, nil
}

func (c *Client) newRawRequest(requestContentType, method, path string, headers map[string]string, queryParams any, body any, out any) error {
	url := fmt.Sprintf("%s%s", c.baseURL, path)

	c.logger.Debug("received request", zap.String("method", method), zap.String("url", url), zap.String("requestContentType", requestContentType), zap.Any("body", body))

	stringBody := ""
	// convert body to string
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			c.logger.Error("error on converting body to string", zap.Error(err))
			return err
		}
		stringBody = string(bodyBytes)
	}

	req, err := http.NewRequest(method, url, strings.NewReader(stringBody))
	if err != nil {
		c.logger.Error("error on getting a new request", zap.Error(err))
		return err
	}
	if method == http.MethodPost && requestContentType != "" {
		req.Header.Set("Content-Type", requestContentType)
	}

	req.Header.Set("Accept", "application/json")

	// include global headers
	if len(c.headers) > 0 {
		for k, v := range c.headers {
			req.Header.Del(k)
			req.Header.Set(k, v)
		}
	}

	// include local headers
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Del(k)
			req.Header.Set(k, v)
		}
	}

	// convert queryParameters
	_queryParameters, err := toMap(queryParams)
	if err != nil {
		c.logger.Error("error on converting queryParameters")
		return err
	}

	if queryParams != nil {
		q := req.URL.Query()
		for k, v := range _queryParameters {
			q.Add(k, fmt.Sprintf("%v", v))
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("error on executing a request", zap.Error(err))
		return err
	}

	c.logger.Debug("response received", zap.String("url", url), zap.String("status", resp.Status))

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error on reading a response body", zap.Error(err))
		return err
	}
	c.logger.Debug("received bytes", zap.String("data", string(respBytes)))

	if resp.StatusCode != http.StatusOK {
		c.logger.Error("request failed", zap.Int("statusCode", resp.StatusCode))
		return fmt.Errorf("failed with status code. [status: %v, statusCode: %v]", resp.Status, resp.StatusCode)
	}

	if out != nil {

		err = json.Unmarshal(respBytes, &out)
		if err != nil {
			c.logger.Error("error on converting to target type", zap.Error(err))
			return err
		}
	}

	return nil
}

func (c *Client) Get(api string, headers map[string]string, queryParams any, out any) error {
	return c.newRawRequest(requestContentType, http.MethodGet, api, headers, queryParams, nil, out)
}

func (c *Client) Post(api string, headers map[string]string, queryParams any, body any, out any) error {
	return c.newRawRequest(requestContentType, http.MethodPost, api, headers, queryParams, body, out)
}

func (c *Client) Delete(api string, headers map[string]string, queryParams any, out any) error {
	return c.newRawRequest(requestContentType, http.MethodDelete, api, headers, queryParams, nil, out)
}
