package main

import (
	"os"
	"strings"
)

type Args struct {
	output string
	stock  string
	types  string
}

func ParseArgs() *Args {
	args := Args{}
	osArgs := os.Args
	if len(osArgs) < 2 {
		return &args
	}

	flag := ""
	for i := range osArgs[1:] {
		arg := osArgs[i]
		if strings.HasPrefix(arg, "-") {
			flag = arg
			continue
		}
		if flag == "" {
			continue
		}
		switch flag {
		case "-o", "-out", "-output":
			args.output = arg
		case "-s", "-stock":
			args.stock = arg
		case "-t", "-type", "-types":
			args.types = arg
		}
		flag = ""
	}
	return &args
}

func (args *Args) IsHelp() bool {
	cmd := args.Command()
	return cmd == "" || cmd == "help" || cmd == "-h"
}

func (args *Args) Command() string {
	xs := os.Args
	if len(xs) < 2 {
		return ""
	}
	return xs[1]
}

func (args *Args) Endpoint() string {
	xs := os.Args
	if len(xs) == 0 {
		return ""
	}
	return xs[len(xs)-1]
}

func (args *Args) HasEndpoint() bool {
	return strings.HasPrefix(args.Endpoint(), "http")
}
