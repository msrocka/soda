package main

import (
	"os"
	"strings"

	"github.com/msrocka/ilcd"
	"github.com/msrocka/soda"
)

type Args struct {
	output string
	stock  string
	types  string
	format string
}

func ParseArgs() *Args {
	args := Args{}
	osArgs := os.Args
	if len(osArgs) < 2 {
		return &args
	}

	flag := ""
	for i := range osArgs[1:] {
		arg := osArgs[i]
		if strings.HasPrefix(arg, "-") {
			flag = arg
			continue
		}
		if flag == "" {
			continue
		}
		switch flag {
		case "-o", "-out", "-output":
			args.output = arg
		case "-s", "-stock":
			args.stock = arg
		case "-t", "-type", "-types":
			args.types = arg
		case "-f", "-format":
			args.format = arg
		}
		flag = ""
	}
	return &args
}

func (args *Args) IsHelp() bool {
	cmd := args.Command()
	return cmd == "" || cmd == "help" || cmd == "-h"
}

func (args *Args) Command() string {
	xs := os.Args
	if len(xs) < 2 {
		return ""
	}
	return xs[1]
}

func (args *Args) Endpoint() string {
	xs := os.Args
	if len(xs) == 0 {
		return ""
	}
	return xs[len(xs)-1]
}

func (args *Args) HasEndpoint() bool {
	return strings.HasPrefix(args.Endpoint(), "http")
}

func (args *Args) Client() *soda.Client {
	client := soda.NewClient(args.Endpoint())
	if args.stock != "" {
		client.WithDataStock(args.stock)
	}
	return client
}

func (args *Args) Types() *TypeFilter {
	filter := make(map[ilcd.DataSetType]bool)
	for _, s := range strings.Split(args.types, ",") {
		t := strings.TrimSpace(strings.ToLower(s))
		switch t {
		case "model", "models":
			filter[ilcd.ModelDataSet] = true
		case "method", "methods":
			filter[ilcd.MethodDataSet] = true
		case "process", "processes", "epd", "epds":
			filter[ilcd.ProcessDataSet] = true
		case "flow", "flows":
			filter[ilcd.FlowDataSet] = true
		case "flow-property", "flow-properties":
			filter[ilcd.FlowPropertyDataSet] = true
		case "unit-group", "unit-groups":
			filter[ilcd.UnitGroupDataSet] = true
		}
	}
	return &TypeFilter{filter}
}

type TypeFilter struct {
	filter map[ilcd.DataSetType]bool
}

func (tf *TypeFilter) DirectList() []ilcd.DataSetType {
	if len(tf.filter) == 0 {
		return tf.all()
	}
	filtered := make([]ilcd.DataSetType, 0)
	for _, t := range tf.all() {
		if tf.filter[t] {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

func (tf *TypeFilter) TransitiveList() []ilcd.DataSetType {
	if len(tf.filter) == 0 {
		return tf.all()
	}
	m := make(map[ilcd.DataSetType]bool)
	for k, v := range tf.filter {
		m[k] = v
	}

	if m[ilcd.ModelDataSet] {
		m[ilcd.ProcessDataSet] = true
	}
	if m[ilcd.ProcessDataSet] || m[ilcd.MethodDataSet] {
		m[ilcd.FlowDataSet] = true
		m[ilcd.SourceDataSet] = true
		m[ilcd.ContactDataSet] = true
	}
	if m[ilcd.FlowDataSet] {
		m[ilcd.FlowPropertyDataSet] = true
	}
	if m[ilcd.FlowPropertyDataSet] {
		m[ilcd.UnitGroupDataSet] = true
	}

	filtered := make([]ilcd.DataSetType, 0)
	for _, t := range tf.all() {
		if m[t] {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

func (tf *TypeFilter) all() []ilcd.DataSetType {
	return []ilcd.DataSetType{
		ilcd.ContactDataSet,
		ilcd.FlowDataSet,
		ilcd.FlowPropertyDataSet,
		ilcd.MethodDataSet,
		ilcd.ModelDataSet,
		ilcd.ProcessDataSet,
		ilcd.SourceDataSet,
		ilcd.UnitGroupDataSet,
	}
}
