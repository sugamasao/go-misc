package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

type commandOption struct {
	opt1 string
	opt2 bool
}

func parseSubCommandOptions(arguments []string) (string, commandOption) {
	options := &commandOption{}
	subcommand := flag.NewFlagSet(arguments[1], flag.ExitOnError)

	subcommand.StringVar(&options.opt1, "message", "", "文字列")
	subcommand.BoolVar(&options.opt2, "exec", false, "真偽値")
	subcommand.Parse(arguments[2:])

	return subcommand.Name(), *options
}

func usage(command string) {
	var mesage = `Usage %s <SubCommand> [Options]
	SubCommands:
		foo FooSubcommand
		bar BarSubcommand
	`
	fmt.Printf(mesage, path.Base(command))
}

func fooCommand(option commandOption) {
	fmt.Println(option)
}

func barCommand(option commandOption) {
	fmt.Println(option)
}

func main() {
	subcommand, options := parseSubCommandOptions(os.Args)

	switch subcommand {
	case "foo":
		fooCommand(options)
	case "bar":
		barCommand(options)
	default:
		usage(os.Args[0])
		os.Exit(1)
	}
}
