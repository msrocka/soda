package soda

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/msrocka/ilcd"
)

type Client struct {
	endpoint string
	stock    string
}

func NewClient(endpoint string) *Client {
	url := strings.TrimSuffix(endpoint, "/")
	if !strings.HasSuffix(url, "/resource") {
		url += "/resource"
	}
	return &Client{endpoint: url}
}

func (client *Client) WithDataStock(stock string) error {
	stocks, err := client.GetDataStocks()
	if err != nil {
		return err
	}
	for i := range stocks.DataStocks {
		s := &stocks.DataStocks[i]
		if s.ShortName == stock || s.UUID == stock {
			client.stock = s.UUID
			return nil
		}
	}
	return fmt.Errorf("could not find data stock: %s", stock)
}

func (client *Client) getRaw(path string) ([]byte, error) {
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
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (client *Client) get(path string, inst any) error {
	bytes, err := client.getRaw(path)
	if err != nil {
		return err
	}
	return xml.Unmarshal(bytes, inst)
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

func (client *Client) GetList(t ilcd.DataSetType) (*DataSetList, error) {
	return client.GetListFor(t, DefaultQuery())
}

func (client *Client) GetListFor(t ilcd.DataSetType, q Query) (*DataSetList, error) {
	path := client.pathOf(t) + q.encode()
	list := DataSetList{}
	if err := client.get(path, &list); err != nil {
		return nil, err
	} else {
		return &list, nil
	}
}

func (client *Client) GetDataSet(t ilcd.DataSetType, id string) ([]byte, error) {
	return client.getRaw(client.pathOf(t) + "/" + id + "?format=XML")
}

func (client *Client) GetMethod(id string) (*ilcd.Method, error) {
	data, err := client.GetDataSet(ilcd.MethodDataSet, id)
	if err != nil {
		return nil, err
	}
	return ilcd.ReadMethod(data)
}

func (client *Client) GetMethods() (*DataSetList, error) {
	return client.GetMethodsFor(DefaultQuery())
}

func (client *Client) GetMethodsFor(q Query) (*DataSetList, error) {
	return client.GetListFor(ilcd.MethodDataSet, q)
}

func (client *Client) GetModel(id string) (*ilcd.Model, error) {
	data, err := client.GetDataSet(ilcd.ModelDataSet, id)
	if err != nil {
		return nil, err
	}
	return ilcd.ReadModel(data)
}

func (client *Client) GetModels() (*DataSetList, error) {
	return client.GetModelsFor(DefaultQuery())
}

func (client *Client) GetModelsFor(q Query) (*DataSetList, error) {
	return client.GetListFor(ilcd.ModelDataSet, q)
}

func (client *Client) GetProcess(id string) (*ilcd.Process, error) {
	data, err := client.GetDataSet(ilcd.ProcessDataSet, id)
	if err != nil {
		return nil, err
	}
	return ilcd.ReadProcess(data)
}

func (client *Client) GetProcesses() (*DataSetList, error) {
	return client.GetProcessesFor(DefaultQuery())
}

func (client *Client) GetProcessesFor(q Query) (*DataSetList, error) {
	return client.GetListFor(ilcd.ProcessDataSet, q)
}

func (client *Client) GetFlow(id string) (*ilcd.Flow, error) {
	data, err := client.GetDataSet(ilcd.FlowDataSet, id)
	if err != nil {
		return nil, err
	}
	return ilcd.ReadFlow(data)
}

func (client *Client) GetFlows() (*DataSetList, error) {
	return client.GetFlowsFor(DefaultQuery())
}

func (client *Client) GetFlowsFor(q Query) (*DataSetList, error) {
	return client.GetListFor(ilcd.FlowDataSet, q)
}

func (client *Client) GetFlowProperty(id string) (*ilcd.FlowProperty, error) {
	data, err := client.GetDataSet(ilcd.FlowPropertyDataSet, id)
	if err != nil {
		return nil, err
	}
	return ilcd.ReadFlowProperty(data)
}

func (client *Client) GetFlowProperties() (*DataSetList, error) {
	return client.GetFlowPropertiesFor(DefaultQuery())
}

func (client *Client) GetFlowPropertiesFor(q Query) (*DataSetList, error) {
	return client.GetListFor(ilcd.FlowPropertyDataSet, q)
}

func (client *Client) GetUnitGroup(id string) (*ilcd.UnitGroup, error) {
	data, err := client.GetDataSet(ilcd.UnitGroupDataSet, id)
	if err != nil {
		return nil, err
	}
	return ilcd.ReadUnitGroup(data)
}

func (client *Client) GetUnitGroups() (*DataSetList, error) {
	return client.GetUnitGroupsFor(DefaultQuery())
}

func (client *Client) GetUnitGroupsFor(q Query) (*DataSetList, error) {
	return client.GetListFor(ilcd.UnitGroupDataSet, q)
}

func (client *Client) GetContact(id string) (*ilcd.Contact, error) {
	data, err := client.GetDataSet(ilcd.ContactDataSet, id)
	if err != nil {
		return nil, err
	}
	return ilcd.ReadContact(data)
}

func (client *Client) GetContacts() (*DataSetList, error) {
	return client.GetContactsFor(DefaultQuery())
}

func (client *Client) GetContactsFor(q Query) (*DataSetList, error) {
	return client.GetListFor(ilcd.ContactDataSet, q)
}

func (client *Client) GetSource(id string) (*ilcd.Source, error) {
	data, err := client.GetDataSet(ilcd.SourceDataSet, id)
	if err != nil {
		return nil, err
	}
	return ilcd.ReadSource(data)
}

func (client *Client) GetSources() (*DataSetList, error) {
	return client.GetSourcesFor(DefaultQuery())
}

func (client *Client) GetSourcesFor(q Query) (*DataSetList, error) {
	return client.GetListFor(ilcd.SourceDataSet, q)
}

func (client *Client) pathOf(t ilcd.DataSetType) string {
	var suffix string
	switch t {
	case ilcd.ContactDataSet:
		suffix = "/contacts"
	case ilcd.FlowDataSet:
		suffix = "/flows"
	case ilcd.FlowPropertyDataSet:
		suffix = "/flowproperties"
	case ilcd.MethodDataSet:
		suffix = "/lciamethods"
	case ilcd.ModelDataSet:
		suffix = "/lifecyclemodels"
	case ilcd.ProcessDataSet:
		suffix = "/processes"
	case ilcd.SourceDataSet:
		suffix = "/sources"
	case ilcd.UnitGroupDataSet:
		suffix = "/unitgroups"
	default:
		suffix = "/unknown"
	}
	if client.stock == "" {
		return suffix
	} else {
		return "/datastocks/" + client.stock + suffix
	}
}
