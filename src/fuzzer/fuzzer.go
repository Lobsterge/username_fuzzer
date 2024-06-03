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
		usernames := generateUsernames(name, surname)
		for _, username := range usernames {
			outputFile.WriteString(username+"\n")
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
	
	permutations := generateAllPermutations(commonNames)
	for _, username := range permutations {
		outputFile.WriteString(username+"\n")
	}
}

func generateUsernames(name, surname string) []string {
	var usernames []string
	firstLetterName := string(name[0])
	firstLetterSurname := string(surname[0])

	usernames = append(usernames, name+surname)
	usernames = append(usernames, surname+name)
	usernames = append(usernames, name+"."+surname)
	usernames = append(usernames, surname+"."+name)
	usernames = append(usernames, name+"-"+surname)
	usernames = append(usernames, surname+"-"+name)
	usernames = append(usernames, name+firstLetterSurname)
	usernames = append(usernames, surname+firstLetterName)
	usernames = append(usernames, firstLetterName+surname)
	usernames = append(usernames, firstLetterSurname+name)

	return usernames
}

func generateAllPermutations(names []string) []string {
	var permutations []string
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
			permutations = append(permutations, generateUsernames(name1, surname2)...)
			permutations = append(permutations, generateUsernames(name2, surname1)...)
		}
	}
	return permutations
}