package fuzzer

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"github.com/Lobsterge/username_fuzzer/src/settings"
)

func FuzzFromFile(args *settings.Settings) {
	inputFile, err := os.Open(args.InputFilePath)

	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", args.InputFilePath, err)
		os.Exit(1)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(args.OutputFilePath)

	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", args.OutputFilePath, err)
		os.Exit(1)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ".")
		if len(parts) != 2 {
			//fmt.Printf("Invalid format: %s\n", line)
			continue
		}
		name := parts[0]
		surname := parts[1]
		usernames := generateUsernames(args, name, surname)
		for _, username := range usernames {
			if args.CaseSensitive {
				outputFile.WriteString(username+"\n")
			} else {
				outputFile.WriteString(strings.ToLower(username+"\n"))
			}
		}
	}
}

func FuzzFromCommon(args *settings.Settings) {
	inputFile, err := os.Open(args.InputFilePath)

	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", args.InputFilePath, err)
		os.Exit(1)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(args.OutputFilePath)

	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", args.OutputFilePath, err)
		os.Exit(1)
	}
	defer outputFile.Close()

	var commonNames []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		commonNames = append(commonNames, line)
	}
	
	permutations := generateAllPermutations(args, commonNames)
	for _, username := range permutations {
		if args.CaseSensitive {
			outputFile.WriteString(username+"\n")
		} else {
			outputFile.WriteString(strings.ToLower(username+"\n"))
		}
	}
}

func generateUsernames(args *settings.Settings, name, surname string) []string {
	var usernames []string
	firstLetterName := string(name[0])
	firstLetterSurname := string(surname[0])
	uppercaseName := strings.ToUpper(name)
    uppercaseSurname := strings.ToUpper(surname)
    lowercaseSurname := strings.ToLower(surname)

	usernames = append(usernames, name+surname)
	usernames = append(usernames, surname+name)
	usernames = append(usernames, name+"."+surname)
	usernames = append(usernames, surname+"."+name)
	usernames = append(usernames, name+"_"+surname)
	usernames = append(usernames, surname+"_"+name)
	usernames = append(usernames, name+"-"+surname)
	usernames = append(usernames, surname+"-"+name)
	usernames = append(usernames, name+firstLetterSurname)
	usernames = append(usernames, surname+firstLetterName)
	usernames = append(usernames, firstLetterName+surname)
	usernames = append(usernames, firstLetterSurname+name)
	usernames = append(usernames, name + firstLetterName + firstLetterSurname + surname)

	if args.CaseSensitive {
		usernames = append(usernames, uppercaseName + uppercaseSurname)
		usernames = append(usernames, uppercaseName + lowercaseSurname)
		usernames = append(usernames, strings.Title(name) + " " + strings.Title(surname))
		usernames = append(usernames, strings.Title(name) + " " + surname)
		usernames = append(usernames, strings.Title(name) + strings.Title(surname))
	}

	for i := 2; i < 5; i++ {
		if len(name)>i && len(surname)>i {
			usernames = append(usernames, name[:i] + surname[:i])
			usernames = append(usernames, surname[:i] + name[:i])
		}
	}

	return usernames
}

func generateAllPermutations(args *settings.Settings, names []string) []string {
    var permutations []string
    addedPermutations := make(map[string]bool)

    for _, fullName1 := range names {
        parts1 := strings.Split(fullName1, ".")
        if len(parts1) != 2 {
            //fmt.Printf("Invalid format: %s\n", fullName1)
            continue
        }
        name1, surname1 := parts1[0], parts1[1]

        for _, fullName2 := range names {
            parts2 := strings.Split(fullName2, ".")
            if len(parts2) != 2 {
                //fmt.Printf("Invalid format: %s\n", fullName2)
                continue
            }
            name2, surname2 := parts2[0], parts2[1]

            perm1 := generateUsernames(args, name1, surname2)
            for _, p := range perm1 {
                if !addedPermutations[p] {
                    permutations = append(permutations, p)
                    addedPermutations[p] = true
                }
            }

            perm2 := generateUsernames(args, name2, surname1)
            for _, p := range perm2 {
                if !addedPermutations[p] {
                    permutations = append(permutations, p)
                    addedPermutations[p] = true
                }
            }
        }
    }
    return permutations
}