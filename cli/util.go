package main

import (
	"github.com/msrocka/ilcd"
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
