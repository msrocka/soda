package soda

import (
	"encoding/xml"

	"github.com/msrocka/ilcd"
)

type DataStockList struct {
	XMLName    xml.Name    `xml:"dataStockList"`
	DataStocks []DataStock `xml:"dataStock"`
}

type DataStock struct {
	XMLName   xml.Name        `xml:"dataStock"`
	IsRoot    bool            `xml:"root,attr"`
	UUID      string          `xml:"uuid"`
	ShortName string          `xml:"shortName"`
	Name      ilcd.LangString `xml:"name"`
}

type DataSetList struct {
	XMLName    xml.Name `xml:"dataSetList"`
	TotalSize  int      `xml:"totalSize,attr"`
	StartIndex int      `xml:"startIndex,attr"`
	PageSize   int      `xml:"pageSize,attr"`

	Models         []ModelInfo        `xml:"lifeCycleModel"`
	Methods        []MethodInfo       `xml:"LCIAMethod"`
	Processes      []ProcessInfo      `xml:"process"`
	Flows          []FlowInfo         `xml:"flow"`
	FlowProperties []FlowPropertyInfo `xml:"flowProperty"`
	UnitGroups     []UnitGroupInfo    `xml:"unitGroup"`
	Contacts       []ContactInfo      `xml:"contact"`
	Sources        []SourceInfo       `xml:"source"`
}

func (list *DataSetList) Size() int {
	return len(list.Models) +
		len(list.Methods) +
		len(list.Processes) +
		len(list.Flows) +
		len(list.FlowProperties) +
		len(list.UnitGroups) +
		len(list.Contacts) +
		len(list.Sources)
}

func (list *DataSetList) IsEmpty() bool {
	return list.Size() == 0
}

func (list *DataSetList) HasMorePages() bool {
	return (list.StartIndex + list.PageSize) < list.TotalSize
}

type DataSetInfo struct {
	UUID           string              `xml:"uuid"`
	Version        string              `xml:"dataSetVersion"`
	Name           ilcd.LangString     `xml:"name"`
	Classification ilcd.Classification `xml:"classification"`
}

type ModelInfo struct {
	DataSetInfo
}

type MethodInfo struct {
	DataSetInfo
}

type ProcessInfo struct {
	DataSetInfo
}

type FlowInfo struct {
	DataSetInfo
}

type FlowPropertyInfo struct {
	DataSetInfo
}

type UnitGroupInfo struct {
	DataSetInfo
}

type ContactInfo struct {
	DataSetInfo
}

type SourceInfo struct {
	DataSetInfo
}
