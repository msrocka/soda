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
	infos := make([]*soda.DataSetInfo, 0)
	client.EachInfo(t, func(info *soda.DataSetInfo) {
		infos = append(infos, info)
	})
	return infos
}
