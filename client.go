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

func (client *Client) get(path string, inst any) error {
	url := client.endpoint + path
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return xml.Unmarshal(body, inst)
}

func (client *Client) GetDataStocks() (*DataStockList, error) {
	stocks := DataStockList{}
	if err := client.get("/datastocks", &stocks); err != nil {
		return nil, err
	} else {
		return &stocks, nil
	}
}

type Query struct {
	StartIndex   int
	PageSize     int
	Search       bool
	Distributed  bool
	Name         string
	Description  string
	ClassId      string
	Lang         string
	LangFallback bool
	AllVersions  bool
	CountOnly    bool
	format       string
}

func DefaultQuery() Query {
	return Query{
		StartIndex: 0,
		PageSize:   500,
	}
}

func (client *Client) GetMethods() (*DataSetList, error) {
	return client.GetMethodsFor(DefaultQuery())
}

func (client *Client) GetMethodsFor(q Query) (*DataSetList, error) {
	path := "/lciamethods"
	// todo: apply query parameters
	list := DataSetList{}
	if err := client.get(path, &list); err != nil {
		return nil, err
	} else {
		return &list, nil
	}
}
