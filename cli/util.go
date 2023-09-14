package main

import (
	"github.com/msrocka/ilcd"
	"github.com/msrocka/soda"
)

func DataSetLabelOf(t ilcd.DataSetType) string {
	switch t {
	case ilcd.Asset:
		return "Assets"
	case ilcd.ContactDataSet:
		return "Contacts"
	case ilcd.ExternalDoc:
		return "External docs"
	case ilcd.FlowDataSet:
		return "Flows"
	case ilcd.FlowPropertyDataSet:
		return "Flow properties"
	case ilcd.MethodDataSet:
		return "Methods"
	case ilcd.ModelDataSet:
		return "Models"
	case ilcd.ProcessDataSet:
		return "Processes"
	case ilcd.SourceDataSet:
		return "Sources"
	case ilcd.UnitGroupDataSet:
		return "Unit groups"
	default:
		return "Unknown data sets"
	}
}

func ScanPage(page *soda.DataSetList, f func(info *soda.DataSetInfo)) {
	for i := range page.Models {
		f(&page.Models[i].DataSetInfo)
	}
	for i := range page.Methods {
		f(&page.Methods[i].DataSetInfo)
	}
	for i := range page.Processes {
		f(&page.Processes[i].DataSetInfo)
	}
	for i := range page.Flows {
		f(&page.Flows[i].DataSetInfo)
	}
	for i := range page.FlowProperties {
		f(&page.FlowProperties[i].DataSetInfo)
	}
	for i := range page.UnitGroups {
		f(&page.UnitGroups[i].DataSetInfo)
	}
	for i := range page.Sources {
		f(&page.Sources[i].DataSetInfo)
	}
	for i := range page.Contacts {
		f(&page.Contacts[i].DataSetInfo)
	}
}
