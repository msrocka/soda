package soda

func scanPage(page *DataSetList, f func(info *DataSetInfo) error) error {

	for i := range page.Models {
		if err := f(&page.Models[i].DataSetInfo); err != nil {
			return err
		}
	}

	for i := range page.Methods {
		if err := f(&page.Methods[i].DataSetInfo); err != nil {
			return err
		}
	}

	for i := range page.Processes {
		if err := f(&page.Processes[i].DataSetInfo); err != nil {
			return err
		}
	}

	for i := range page.Flows {
		if err := f(&page.Flows[i].DataSetInfo); err != nil {
			return err
		}
	}

	for i := range page.FlowProperties {
		if err := f(&page.FlowProperties[i].DataSetInfo); err != nil {
			return err
		}
	}

	for i := range page.UnitGroups {
		if err := f(&page.UnitGroups[i].DataSetInfo); err != nil {
			return err
		}
	}

	for i := range page.Sources {
		if err := f(&page.Sources[i].DataSetInfo); err != nil {
			return err
		}
	}

	for i := range page.Contacts {
		if err := f(&page.Contacts[i].DataSetInfo); err != nil {
			return err
		}
	}

	return nil
}
