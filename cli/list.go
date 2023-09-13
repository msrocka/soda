package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/msrocka/soda"
)

func ListStocks(args *Args) {
	client := soda.NewClient(args.Endpoint())
	stocks, err := client.GetDataStocks()
	Check(err, "failed to get data stocks")
	if len(stocks.DataStocks) == 0 {
		fmt.Println("no data stocks found")
		return
	}
	template := "| %-36s | %-10s | %s\n"
	_, err = fmt.Printf(template, "UUID", "Is root?", "Name")
	Check(err, "invalid format")
	for _, s := range stocks.DataStocks {
		isRoot := "false"
		if s.IsRoot {
			isRoot = "true"
		}
		_, err = fmt.Printf(template, s.UUID, isRoot, s.ShortName)
		Check(err, "invalid format")
	}
}

func ListDataSets(args *Args) {
	client := soda.NewClient(args.Endpoint())
	q := soda.DefaultQuery()
	q.PageSize = 1000
	p := printerOf(args)
	for {
		page, err := client.GetProcessesFor(q)
		Check(err, "failed to get data sets")

		template := "| %-36s | %-10s | %s\n"
		_, err = fmt.Printf(template, "UUID", "Version", "Name")
		p.header("UUID", "Version", "Name")
		printPage(p, page)
		Check(err, "invalid format")
		for _, ds := range page.Processes {
			fmt.Printf(template, ds.UUID, ds.Version, ds.Name.Default())
		}
		if !page.HasMorePages() {
			break
		}
		q = q.NextPage()
	}
	q.NextPage()
}

type printer interface {
	header(...string)
	dataSet(*soda.DataSetInfo)
}

func printPage(p printer, page *soda.DataSetList) {
	for i := range page.Models {
		p.dataSet(&page.Models[i].DataSetInfo)
	}
	for i := range page.Methods {
		p.dataSet(&page.Methods[i].DataSetInfo)
	}
	for i := range page.Processes {
		p.dataSet(&page.Processes[i].DataSetInfo)
	}
	for i := range page.Flows {
		p.dataSet(&page.Flows[i].DataSetInfo)
	}
	for i := range page.FlowProperties {
		p.dataSet(&page.FlowProperties[i].DataSetInfo)
	}
	for i := range page.UnitGroups {
		p.dataSet(&page.UnitGroups[i].DataSetInfo)
	}
	for i := range page.Sources {
		p.dataSet(&page.Sources[i].DataSetInfo)
	}
	for i := range page.Contacts {
		p.dataSet(&page.Contacts[i].DataSetInfo)
	}
}

func printerOf(args *Args) printer {
	if args.output != "" {
		file, err := os.Open(args.output)
		Check(err, "failed to open output file")
		writer := csv.NewWriter(file)
		return &csvPrinter{writer}
	}
	return &console{
		template: "| %-36s | %-10s | %s\n",
	}
}

type console struct {
	template string
}

func (c *console) header(header ...string) {
	_, err := fmt.Printf(c.template, header)
	Check(err, "invalid format")
}

func (c *console) dataSet(i *soda.DataSetInfo) {
	name := strings.ReplaceAll(i.Name.Default(), "\n", " ")
	_, err := fmt.Printf(c.template, i.UUID, i.Version, name)
	Check(err, "invalid format")
}

type csvPrinter struct {
	file   os.File
	writer *csv.Writer
}

func (c *csvPrinter) header(header ...string) {
	Check(c.writer.Write(header), "failed to write row")
}

func (c *csvPrinter) dataSet(i *soda.DataSetInfo) {
	err := c.writer.Write([]string{i.UUID, i.Version, i.Name.Default()})
	Check(err, "failed to write row")
}

func (c *csvPrinter) Close() {
	c.file.Close()
}
