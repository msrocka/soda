package soda

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
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

func (client *Client) WithDataStock(idOrName string) error {
	stock, err := client.findDataStock(idOrName)
	if err != nil {
		return err
	}
	client.stock = stock.UUID
	return nil
}

func (client *Client) getRaw(path string) ([]byte, error) {
	var data []byte
	err := client.handleBody(path, func(body io.Reader) error {
		if bytes, err := io.ReadAll(body); err != nil {
			return err
		} else {
			data = bytes
			return nil
		}
	})
	return data, err
}

func (client *Client) handleBody(path string, f func(io.Reader) error) error {
	url := client.endpoint + path
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf(
			"request GET %s failed: %d %s", url, res.StatusCode, res.Status)
	}
	f(res.Body)
	return res.Body.Close()
}

func (client *Client) get(path string, inst any) error {
	bytes, err := client.getRaw(path)
	if err != nil {
		return err
	}
	return xml.Unmarshal(bytes, inst)
}

func (client *Client) GetList(t ilcd.DataSetType) (*DataSetList, error) {
	return client.GetListFor(t, DefaultQuery())
}

func (client *Client) GetListFor(t ilcd.DataSetType, q Query) (*DataSetList, error) {
	path := client.pathOf(t) + q.String()
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

func (client *Client) EachInfo(t ilcd.DataSetType, f func(*DataSetInfo) error) error {
	q := DefaultQuery()
	q.PageSize = 5000
	for {
		page, err := client.GetListFor(t, q)
		if err != nil {
			return err
		}
		if page.IsEmpty() {
			break
		}
		if err := scanPage(page, f); err != nil {
			return err
		}
		if !page.HasMorePages() {
			break
		}
		q = q.NextPage()
	}
	return nil
}

func (client *Client) EachDataSet(
	t ilcd.DataSetType, f func(*DataSetInfo, []byte) error) error {
	return client.EachInfo(t, func(info *DataSetInfo) error {
		data, err := client.GetDataSet(t, info.UUID)
		if err != nil {
			return err
		}
		return f(info, data)
	})
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
