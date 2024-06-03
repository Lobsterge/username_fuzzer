package settings

import (
	"fmt"
	"os"
)

type Settings struct {
	Command    string
	InputFilePath string
	OutputFilePath string
	Help bool
}

func New() *Settings {
	return &Settings{
		Command: "",
		InputFilePath: "",
		OutputFilePath: "",
		Help: false,
	}
}

//Verifies file paths are correct
func (settings *Settings) Verify() (err error) {

	if err = settings.verifyOutputPath(); err != nil {
		return err
	}

	if err = settings.verifyInputPath(); err != nil {
		return err
	}
	
	return err
}

func (settings *Settings) verifyOutputPath() (err error) {

	info, err := os.Stat(settings.OutputFilePath)

	if os.IsNotExist(err) {
		file, err2 := os.Create(settings.OutputFilePath)
   		if err2 != nil {
			return fmt.Errorf("Could not create output file %s", settings.OutputFilePath)
   		}
   		defer file.Close()
		return err2
	}

	if info.Mode().IsDir() {
		return fmt.Errorf("Output file %s is a directory", settings.OutputFilePath)
	}

	return err
}

func (settings *Settings) verifyInputPath() (err error) {

	info, err := os.Stat(settings.InputFilePath)

	if os.IsNotExist(err) {
		return fmt.Errorf("Input file %s does not exist", settings.InputFilePath)
	}

	if info.Mode().IsDir() {
		return fmt.Errorf("Input file %s is a directory", settings.InputFilePath)
	}

	return err
}