package main

import (
	"flag"
	"strings"
)

type Args struct {
	output string
	stock  string
	types  string
	help   bool
}

func ParseArgs() *Args {
	args := &Args{}
	flag.StringVar(&args.output, "o", "",
		"The path where the output should be written to.")
	flag.StringVar(&args.stock, "ds", "",
		"The optional data stock that should be used.")
	flag.StringVar(&args.types, "types", "",
		"The optional data set filter; e.g. '-types processes,flows'")
	flag.BoolVar(&args.help, "h", false, "Prints this help.")
	flag.Parse()
	return args
}

func (args *Args) IsHelp() bool {
	if args.help {
		return true
	}
	cmd := args.Command()
	return cmd == "" || cmd == "help"
}

func (args *Args) Command() string {
	xs := flag.Args()
	if len(xs) == 0 {
		return ""
	}
	return xs[0]
}

func (args *Args) Endpoint() string {
	xs := flag.Args()
	if len(xs) == 0 {
		return ""
	}
	return xs[len(xs)-1]
}

func (args *Args) HasEndpoint() bool {
	return strings.HasPrefix(args.Endpoint(), "http")
}
