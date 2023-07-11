package soda

import (
	"encoding/xml"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	endpoint string
}

func NewClient(endpoint string) *Client {
	url := strings.TrimSuffix(endpoint, "/")
	if !strings.HasSuffix(url, "/resource") {
		url += "/resource"
	}
	return &Client{url}
}

func (client *Client) get(path string) ([]byte, error) {
	url := client.endpoint + path
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func (client *Client) GetDataStocks() (*DataStockList, error) {
	body, err := client.get("/datastocks")
	if err != nil {
		return nil, err
	}
	stocks := DataStockList{}
	if err := xml.Unmarshal(body, &stocks); err != nil {
		return nil, err
	}
	return &stocks, nil
}
