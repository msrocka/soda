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
}

type MethodInfo struct {
}

type FlowInfo struct {
}
