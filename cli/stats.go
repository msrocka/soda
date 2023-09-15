package main

import (
	"fmt"
	"log"

	"github.com/msrocka/soda"
)

func Stats(args *Args) {
	stats := collectStats(args)
	fmt.Println()
	fmt.Printf("| %-16s | %-5s |\n", "Data set type", "Count")
	for _, stat := range stats {
		fmt.Printf("| %-16s | % 5d |\n", stat.label, stat.count)
	}
	fmt.Println()
}

type stat struct {
	label string
	count int
}

func collectStats(args *Args) []stat {
	log.Println("collect statistics from server")
	types := args.Types().DirectList()
	stats := make([]stat, 0, len(types))
	client := args.Client()
	q := soda.DefaultQuery()
	q.PageSize = 1
	for _, t := range types {
		page, err := client.GetListFor(t, q)
		var count int
		if err != nil {
			log.Println("failed to get statistics for type:", t)
			count = 0
		} else {
			count = page.TotalSize
		}
		stats = append(stats, stat{DataSetLabelOf(t), count})
	}
	return stats
}
