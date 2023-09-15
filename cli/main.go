package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/msrocka/ilcd"
	"github.com/msrocka/soda"
)

func main() {
	args := ParseArgs()
	if args.IsHelp() {
		printHelp()
		os.Exit(0)
	}

	if !args.HasEndpoint() {
		fmt.Println("error: invalid arguments,",
			"the last argument needs to be a valid endpoint; see 'help'")
		os.Exit(1)
	}

	switch args.Command() {
	case "fetch":
		FetchDataSets(args)
	case "list":
		ListDataSets(args)
	case "stocks":
		ListStocks(args)
	case "stats":
		Stats(args)
	case "export":
		ExportStock(args)
	}
}

func Check(err error, msg string) {
	if err != nil {
		fmt.Println("error:", msg)
		fmt.Println("Error details:", err)
		os.Exit(1)
	}
}

func printHelp() {
	progPath := strings.Split(os.Args[0], string(os.PathSeparator))
	prog := progPath[len(progPath)-1]
	fmt.Println("Usage of", strings.TrimSuffix(prog, ".exe"))
	// TODO: document commands
	flag.PrintDefaults()
}

func FetchDataSets(args *Args) {
	if args.output == "" {
		fmt.Println("error: no output file provided; usage: -o path/to/file.zip")
		return
	}

	log.Println("download data sets to package", args.output)
	zip, err := ilcd.NewZipWriter(args.output)
	Check(err, "failed to create zip file")
	defer zip.Close()
	client := args.Client()

	for _, t := range args.Types().TransitiveList() {
		label := DataSetLabelOf(t)
		log.Println("download:", label)
		i := 0
		err = client.EachDataSet(t, func(info *soda.DataSetInfo, data []byte) error {
			i += 1
			if i%100 == 0 {
				log.Println(label, "- downloaded", i, "data sets")
			}
			path := "ILCD/" + t.Folder() + "/" + info.UUID + "_" + info.Version + ".xml"
			Check(zip.Write(path, data), "failed to write "+path)
			return nil
		})
		if err != nil {
			log.Println(label, "- failed to download data sets:", err)
		} else {
			log.Println(label, "- downloaded", i, "data sets")
		}
	}
}

func ExportStock(args *Args) {
	if args.stock == "" {
		fmt.Println("error: no data stock provided; usage: -s DATA_STOCK")
		return
	}
	if args.output == "" {
		fmt.Println("error: no output file provided; usage: -o PATH/TO/FILE")
		return
	}

	// open file
	file, err := os.Create(args.output)
	Check(err, "failed to create the output file")
	defer file.Close()
	log.Println("export data set", args.stock, "to", args.output)

	// detect format
	asCsv := false
	format := strings.ToLower(strings.TrimSpace(args.format))
	if strings.HasSuffix(format, "csv") {
		asCsv = true
		log.Println("export as CSV file")
	} else if !strings.HasSuffix(format, "zip") {
		log.Println("warning: unknown format", format, "; fall back to zip")
	}

	// make request
	client := args.Client()
	if asCsv {
		err = client.ExportDataStockCSV(args.stock, file)
	} else {
		err = client.ExportDataStock(args.stock, file)
	}
	Check(err, "failed to download data stock")
}
