package brasilapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

var (
	APPLICATION_JSON_UTF8 = "application/json; charset=utf-8"
	BASE_URL              = "https://brasilapi.com.br/api/cep/v1/"
)

type Client struct {
	BaseURL *url.URL
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewClient() *Client {
	baseUrl, err := url.Parse(BASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		BaseURL: baseUrl,
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", APPLICATION_JSON_UTF8)
	req.Header.Set("Accept", APPLICATION_JSON_UTF8)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("erro desconhecido, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(v); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetCepInfo(ctx *context.Context, cep string) (CepInfo, error) {

	url := c.BaseURL.JoinPath(cep).String()
	req, err := http.NewRequestWithContext(*ctx, "GET", url, nil)
	if err != nil {
		return CepInfo{}, err
	}

	var res CepInfo
	if err := c.sendRequest(req, &res); err != nil {
		return CepInfo{}, err
	}

	return res, nil
}
