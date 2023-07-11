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

	Methods []MethodInfo `xml:"LCIAMethod"`
}

type DataSetInfo struct {
	UUID           string              `xml:"uuid"`
	Version        string              `xml:"dataSetVersion"`
	Name           ilcd.LangString     `xml:"name"`
	Classification ilcd.Classification `xml:"classification"`
}

type MethodInfo struct {
	DataSetInfo
}

type FlowInfo struct {
	DataSetInfo
}
