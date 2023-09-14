package soda

func scanPage(page *DataSetList, f func(info *DataSetInfo)) {
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
