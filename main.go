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

	flag.StringVar(&args.Command, "command", "", "file -> generates usernames from input file\nseparate -> generates usernames from a names and surnames files\nitaly -> generates usernames from the most common names in Italy\nworld -> generates usernames from the most common names globally")
	flag.StringVar(&args.Command, "c", "", "-command (shorthand)")

	flag.StringVar(&args.OutputFilePath, "output", "output.txt", "Path of the output file")
	flag.StringVar(&args.OutputFilePath, "o", "output.txt", "-output (shorthand)")

	flag.StringVar(&args.InputFilePath, "input", "", "Path of the input file in the format (name.surname)")
	flag.StringVar(&args.InputFilePath, "i", "", "-input (shorthand)")

	flag.BoolVar(&args.Help, "help", false, "Shows the various command line options")
	flag.BoolVar(&args.Help, "h", false, "-help (shorthand)")

	flag.BoolVar(&args.Permutation, "permutation", false, "Applies a permutation on the list provided by -input")
	flag.BoolVar(&args.Permutation, "p", false, "-permutation (shorthand)")

	flag.BoolVar(&args.CaseSensitive, "case", false, "Make the usernames case-sensitive, if this flag is not checked they will be all lowercase")
	flag.BoolVar(&args.CaseSensitive, "cs", false, "-case (shorthand)")

	flag.StringVar(&args.NamesFilePath, "names", "", "Path of the file containing names for permutation")
	flag.StringVar(&args.NamesFilePath, "n", "", "-names (shorthand)")

	flag.StringVar(&args.SurnamesFilePath, "surnames", "", "Path of the file containing surnames for permutation")
	flag.StringVar(&args.SurnamesFilePath, "s", "", "-surnames (shorthand)")

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

			if args.Permutation {
				fuzzer.FuzzFromCommon(args)
			} else {
				fuzzer.FuzzFromFile(args)
			}
		
		case "separate":
			if args.NamesFilePath == "" || args.SurnamesFilePath == "" {
				fmt.Println("You have to provide two files to use this mode, use -h for more information")
				os.Exit(1)
			}

			if err := args.Verify(); err != nil {
				fmt.Printf("An error has occured: %s\n", err)
				os.Exit(1)
			}

			fuzzer.FuzzFromFiles(args)

		case "italy":
			args.InputFilePath="data/italy.txt"

			if err := args.Verify(); err != nil {
				fmt.Printf("An error has occured: %s\n", err)
				os.Exit(1)
			}

			fuzzer.FuzzFromCommon(args)

		case "world":
			args.InputFilePath="data/world.txt"

			if err := args.Verify(); err != nil {
				fmt.Printf("An error has occured: %s\n", err)
				os.Exit(1)
			}

			fuzzer.FuzzFromCommon(args)
		
		default:
			flag.Usage()
	}
}