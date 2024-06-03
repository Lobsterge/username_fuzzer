package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/Lobsterge/username_fuzzer/src/settings"
)

func parseSettings() (args *settings.Settings) {
	args = settings.New()

	flag.StringVar(&args.Command, "command", "", "Your command")
	flag.StringVar(&args.Command, "c", "", "Your command")

	flag.StringVar(&args.OutputFilePath, "output", "", "Output file")
	flag.StringVar(&args.OutputFilePath, "o", "", "Output file")

	flag.StringVar(&args.InputFilePath, "input", "", "Input file")
	flag.StringVar(&args.InputFilePath, "i", "", "Input file")

	flag.BoolVar(&args.Help, "help", false, "Help")
	flag.BoolVar(&args.Help, "h", false, "Help")

	flag.Parse()

	return args
}

func main() {
	args := parseSettings()

	fmt.Printf("%s\n%s\n%s\n%t\n", args.Command, args.InputFilePath, args.OutputFilePath, args.Help)
	os.Exit(0)
}