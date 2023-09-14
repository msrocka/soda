package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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
	}
}

func Check(err error, msg string) {
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
