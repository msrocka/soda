package soda

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"strconv"
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
	Format       string
}

func (query *Query) NextPage() Query {
	return Query{
		StartIndex:   query.StartIndex + query.PageSize,
		PageSize:     query.PageSize,
		Search:       query.Search,
		Distributed:  query.Distributed,
		Name:         query.Name,
		Description:  query.Description,
		ClassId:      query.ClassId,
		Lang:         query.Lang,
		LangFallback: query.LangFallback,
		AllVersions:  query.AllVersions,
		CountOnly:    query.CountOnly,
		Format:       query.Format,
	}
}

func (query *Query) encode() string {
	params := url.Values{}
	if query.StartIndex != 0 {
		params.Add("startIndex", strconv.Itoa(query.StartIndex))
	}
	if query.PageSize != 500 {
		params.Add("pageSize", strconv.Itoa(query.PageSize))
	}
	if query.Search {
		params.Add("search", "true")
	}
	if query.Distributed {
		params.Add("distributed", "true")
	}
	if query.LangFallback {
		params.Add("langFallback", "true")
	}
	if query.AllVersions {
		params.Add("allVersions", "true")
	}
	if query.CountOnly {
		params.Add("countOnly", "true")
	}
	if len(query.Name) > 0 {
		params.Add("name", query.Name)
	}
	if len(query.Description) > 0 {
		params.Add("description", query.Description)
	}
	if len(query.ClassId) > 0 {
		params.Add("classId", query.ClassId)
	}
	if len(query.Lang) > 0 {
		params.Add("lang", query.Lang)
	}
	if len(params) > 0 {
		return "?" + params.Encode()
	} else {
		return ""
	}
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
	path := "/lciamethods" + q.encode()
	list := DataSetList{}
	if err := client.get(path, &list); err != nil {
		return nil, err
	} else {
		return &list, nil
	}
}

func (client *Client) GetModels() (*DataSetList, error) {
	return client.GetModelsFor(DefaultQuery())
}

func (client *Client) GetModelsFor(q Query) (*DataSetList, error) {
	path := "/lifecyclemodels" + q.encode()
	list := DataSetList{}
	if err := client.get(path, &list); err != nil {
		return nil, err
	} else {
		return &list, nil
	}
}

func (client *Client) GetProcesses() (*DataSetList, error) {
	return client.GetProcessesFor(DefaultQuery())
}

func (client *Client) GetProcessesFor(q Query) (*DataSetList, error) {
	path := "/processes" + q.encode()
	list := DataSetList{}
	if err := client.get(path, &list); err != nil {
		return nil, err
	} else {
		return &list, nil
	}
}

func (client *Client) GetFlows() (*DataSetList, error) {
	return client.GetFlowsFor(DefaultQuery())
}

func (client *Client) GetFlowsFor(q Query) (*DataSetList, error) {
	path := "/flows" + q.encode()
	list := DataSetList{}
	if err := client.get(path, &list); err != nil {
		return nil, err
	} else {
		return &list, nil
	}
}

func (client *Client) GetFlowProperties() (*DataSetList, error) {
	return client.GetFlowPropertiesFor(DefaultQuery())
}

func (client *Client) GetFlowPropertiesFor(q Query) (*DataSetList, error) {
	path := "/flowproperties" + q.encode()
	list := DataSetList{}
	if err := client.get(path, &list); err != nil {
		return nil, err
	} else {
		return &list, nil
	}
}

func (client *Client) GetUnitGroups() (*DataSetList, error) {
	return client.GetUnitGroupsFor(DefaultQuery())
}

func (client *Client) GetUnitGroupsFor(q Query) (*DataSetList, error) {
	path := "/unitgroups" + q.encode()
	list := DataSetList{}
	if err := client.get(path, &list); err != nil {
		return nil, err
	} else {
		return &list, nil
	}
}

func (client *Client) GetContacts() (*DataSetList, error) {
	return client.GetContactsFor(DefaultQuery())
}

func (client *Client) GetContactsFor(q Query) (*DataSetList, error) {
	path := "/contacts" + q.encode()
	list := DataSetList{}
	if err := client.get(path, &list); err != nil {
		return nil, err
	} else {
		return &list, nil
	}
}

func (client *Client) GetSources() (*DataSetList, error) {
	return client.GetSourcesFor(DefaultQuery())
}

func (client *Client) GetSourcesFor(q Query) (*DataSetList, error) {
	path := "/sources" + q.encode()
	list := DataSetList{}
	if err := client.get(path, &list); err != nil {
		return nil, err
	} else {
		return &list, nil
	}
}
