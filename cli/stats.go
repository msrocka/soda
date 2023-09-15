package main

import (
	"fmt"

	"github.com/msrocka/soda"
)

func Stats(args *Args) {
	client := args.Client()
	q := soda.DefaultQuery()
	q.PageSize = 1
	for _, t := range args.Types().DirectList() {
		page, err := client.GetListFor(t, q)
		Check(err, "failed to get stats for "+t.String())
		fmt.Printf("| %-16s | % 5d |\n", DataSetLabelOf(t), page.TotalSize)
	}
}
