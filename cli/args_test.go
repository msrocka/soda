package main

import (
	"testing"

	"slices"

	"github.com/msrocka/ilcd"
)

func TestTypeFi(t *testing.T) {
	args := Args{types: "epds,unit-groups"}
	filter := args.Types()

	direct := filter.DirectList()
	if len(direct) != 2 ||
		!slices.Contains(direct, ilcd.ProcessDataSet) ||
		!slices.Contains(direct, ilcd.UnitGroupDataSet) {
		t.Fatal("unexpected direct list")
	}

	trans := filter.TransitiveList()
	if len(trans) != 6 ||
		!slices.Contains(trans, ilcd.ProcessDataSet) ||
		!slices.Contains(trans, ilcd.FlowDataSet) ||
		!slices.Contains(trans, ilcd.FlowPropertyDataSet) ||
		!slices.Contains(trans, ilcd.UnitGroupDataSet) ||
		!slices.Contains(trans, ilcd.ContactDataSet) ||
		!slices.Contains(trans, ilcd.SourceDataSet) {
		t.Fatal("unexpected trans. list")
	}
}
