package main

import (
	"fmt"

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
	for {
		page, err := client.GetProcessesFor(q)
		Check(err, "failed to get data sets")

		template := "| %-36s | %-10s | %s\n"
		_, err = fmt.Printf(template, "UUID", "Version", "Name")
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
