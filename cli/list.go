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
	p := printerOf(args)
	defer p.close()
	p.section("Data stocks")
	p.header("UUID", "Is root?", "Name")
	for i := range stocks.DataStocks {
		p.stock(&stocks.DataStocks[i])
	}
}

func ListDataSets(args *Args) {
	client := args.Client()
	p := printerOf(args)
	for i, t := range args.Types().DirectList() {
		if i > 0 {
			p.br()
		}
		p.section(DataSetLabelOf(t))
		p.header("UUID", "Version", "Name")
		err := client.EachInfo(t, func(info *soda.DataSetInfo) error {
			p.dataSet(info)
			return nil
		})
		Check(err, "failed to list data set infos")
	}
}

type printer interface {
	header(...string)
	dataSet(*soda.DataSetInfo)
	stock(*soda.DataStock)
	section(string)
	br()
	close()
}

func printerOf(args *Args) printer {
	if args.output != "" {
		file, err := os.Create(args.output)
		Check(err, "failed to open output file")
		writer := csv.NewWriter(file)
		return &csvPrinter{file, writer}
	}
	return &console{
		template: "| %-36s | %-10s | %s\n",
	}
}

type console struct {
	template string
}

func (c *console) header(header ...string) {
	xs := make([]any, len(header))
	for i := range header {
		xs[i] = header[i]
	}
	_, err := fmt.Printf(c.template, xs...)
	Check(err, "invalid format")
}

func (c *console) dataSet(i *soda.DataSetInfo) {
	name := strings.ReplaceAll(i.Name.Default(), "\n", " ")
	_, err := fmt.Printf(c.template, i.UUID, i.Version, name)
	Check(err, "invalid format")
}

func (c *console) stock(s *soda.DataStock) {
	isRoot := "false"
	if s.IsRoot {
		isRoot = "true"
	}
	_, err := fmt.Printf(c.template, s.UUID, isRoot, s.ShortName)
	Check(err, "invalid format")
}

func (c *console) section(header string) {
	fmt.Println(header)
}

func (c *console) br() {
	fmt.Println()
}

func (c *console) close() {
}

type csvPrinter struct {
	file   *os.File
	writer *csv.Writer
}

func (c *csvPrinter) header(header ...string) {
	c.next(header...)
}

func (c *csvPrinter) dataSet(i *soda.DataSetInfo) {
	c.next(i.UUID, i.Version, i.Name.Default())
}

func (c *csvPrinter) stock(s *soda.DataStock) {
	isRoot := "false"
	if s.IsRoot {
		isRoot = "true"
	}
	c.next(s.UUID, isRoot, s.ShortName)
}

func (c *csvPrinter) section(header string) {
	c.next(header)
}

func (c *csvPrinter) br() {
	c.next()
}

func (c *csvPrinter) next(row ...string) {
	Check(c.writer.Write(row), "failed to write row")
}

func (c *csvPrinter) close() {
	c.writer.Flush()
	c.file.Close()
}
