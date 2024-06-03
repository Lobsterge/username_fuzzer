package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/Lobsterge/username_fuzzer/src/settings"
	"github.com/Lobsterge/username_fuzzer/src/fuzzer"
)

func parseSettings() (args *settings.Settings) {
	args = settings.New()

	flag.StringVar(&args.Command, "command", "", "file -> generates usernames from input file\nitaly -> generates usernames from the most common names in Italy\nworld -> generates usernames from the most common names globally")
	flag.StringVar(&args.Command, "c", "", "-command (shorthand)")

	flag.StringVar(&args.OutputFilePath, "output", "output.txt", "Path of the output file")
	flag.StringVar(&args.OutputFilePath, "o", "output.txt", "-output (shorthand)")

	flag.StringVar(&args.InputFilePath, "input", "", "Path of the input file in the format (name.surname)")
	flag.StringVar(&args.InputFilePath, "i", "", "-input (shorthand)")

	flag.BoolVar(&args.Help, "help", false, "Shows the various command line options")
	flag.BoolVar(&args.Help, "h", false, "-help (shorthand)")

	flag.Parse()

	return args
}

func main() {
	args := parseSettings()

	if args.Help {
		flag.Usage()
		os.Exit(0)
	}

	switch args.Command {
		case "file":
			if args.InputFilePath == "" {
				fmt.Println("You have to provide an input file to use this mode, use -h for more information")
				os.Exit(1)
			}

			if err := args.Verify(); err != nil {
				fmt.Printf("An error has occured: %s\n", err)
				os.Exit(1)
			}

			fuzzer.FuzzFromFile(args)

		case "italy":
			args.InputFilePath="data/italy.txt"
			fuzzer.FuzzFromCommon(args)

		case "world":
			args.InputFilePath="data/world.txt"
			fuzzer.FuzzFromCommon(args)
		
		default:
			flag.Usage()
	}
}