package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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
		fmt.Println("fetch ...")
	case "list":
		fmt.Println("list ...")
	case "stocks":
		listStocks(args)
	}
}

func listStocks(args *Args) {
	client := soda.NewClient(args.Endpoint())
	stocks, err := client.GetDataStocks()
	check(err, "failed to get data stocks")
	if len(stocks.DataStocks) == 0 {
		fmt.Println("no data stocks found")
		return
	}
	template := "| %-40s | %-10s | %s"
	fmt.Println(fmt.Sprintln(template, "UUID", "Is root?", "Name"))
	for _, s := range stocks.DataStocks {
		fmt.Println(fmt.Sprintln(template, s.UUID, s.IsRoot, s.Name))
	}
}

func check(err error, msg string) {
	if err != nil {
		fmt.Println("error:", msg, err)
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
