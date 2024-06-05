package settings

import (
	"fmt"
	"os"
)

type Settings struct {
	Command    string
	InputFilePath string
	OutputFilePath string
	NamesFilePath string
    SurnamesFilePath string 
	Help bool
	CaseSensitive bool
	Permutation bool
}

func New() *Settings {
	return &Settings{
		Command: "",
		InputFilePath: "",
		OutputFilePath: "",
		NamesFilePath:   "", 
        SurnamesFilePath: "", 
		Help: false,
		Permutation: false,
		CaseSensitive: false,
	}
}

//Verifies file paths are correct
func (settings *Settings) Verify() (err error) {

	if err = settings.verifyOutputPath(); err != nil {
		return err
	}

	if settings.InputFilePath != "" {
		if err = settings.verifyInputPath(); err != nil {
			return err
		}
	} else {
		if err = settings.verifyNamesPath(); err != nil {
			return err
		}

		if err = settings.verifySurnamesPath(); err != nil {
			return err
		}
	}
	
	return err
}

func (settings *Settings) verifyOutputPath() (err error) {

	info, err := os.Stat(settings.OutputFilePath)

	if os.IsNotExist(err) {
		file, err2 := os.Create(settings.OutputFilePath)
   		if err2 != nil {
			return fmt.Errorf("could not create output file %s", settings.OutputFilePath)
   		}
   		defer file.Close()
		return err2
	}

	if info.Mode().IsDir() {
		return fmt.Errorf("output file %s is a directory", settings.OutputFilePath)
	}

	return err
}

func (settings *Settings) verifyInputPath() (err error) {

	info, err := os.Stat(settings.InputFilePath)

	if os.IsNotExist(err) {
		return fmt.Errorf("input file %s does not exist", settings.InputFilePath)
	}

	if info.Mode().IsDir() {
		return fmt.Errorf("input file %s is a directory", settings.InputFilePath)
	}

	return err
}

func (settings *Settings) verifySurnamesPath() (err error) {

	info, err := os.Stat(settings.SurnamesFilePath)

	if os.IsNotExist(err) {
		return fmt.Errorf("surnames file %s does not exist", settings.SurnamesFilePath)
	}

	if info.Mode().IsDir() {
		return fmt.Errorf("surnames file %s is a directory", settings.SurnamesFilePath)
	}

	return err
}

func (settings *Settings) verifyNamesPath() (err error) {

	info, err := os.Stat(settings.NamesFilePath)

	if os.IsNotExist(err) {
		return fmt.Errorf("names file %s does not exist", settings.NamesFilePath)
	}

	if info.Mode().IsDir() {
		return fmt.Errorf("names file %s is a directory", settings.NamesFilePath)
	}

	return err
}