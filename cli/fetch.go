package main

import (
	"fmt"

	"github.com/msrocka/ilcd"
	"github.com/msrocka/soda"
)

func FetchDataSets(args *Args) {
	if args.output == "" {
		fmt.Println("error: no output file provided; usage: -o path/to/file.zip")
		return
	}

	zip, err := ilcd.NewZipWriter(args.output)
	Check(err, "failed to create zip file")
	defer zip.Close()
	client := args.Client()

	for _, t := range args.Types().TransitiveList() {
		infos := collectInfos(client, t)
		for _, info := range infos {
			data, err := client.GetDataSet(t, info.UUID)
			Check(err, "failed to fetch data set")
			zip.Write("ILCD/"+t.Folder()+"/"+info.UUID+"_"+info.Version+".xml", data)
		}
	}
}

func collectInfos(client *soda.Client, t ilcd.DataSetType) []*soda.DataSetInfo {
	q := soda.DefaultQuery()
	q.PageSize = 1000
	infos := make([]*soda.DataSetInfo, 0)
	println("  fetch descriptors of:", DataSetLabelOf(t))
	fetched := 0
	for {
		page, err := client.GetListFor(t, q)
		Check(err, "failed to get data sets")
		if page.IsEmpty() {
			break
		}
		ScanPage(page, func(info *soda.DataSetInfo) {
			infos = append(infos, info)
		})
		if !page.HasMorePages() {
			break
		}
		fetched += page.Size()
		println("  .. fetched", fetched, "data sets; fetch next page")
		q = q.NextPage()
	}
	return infos
}
