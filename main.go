package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"uzo/util"
	"github.com/spf13/cobra"
)


var(
	File string
	// dest string
)

var rootCmd = &cobra.Command{
	Use:   "uzo command to unzip and open with code",
	Short: "Unzipping application",
	Args: func(cmd *cobra.Command, args []string) error {
		if File == "" && len(args) < 1 {
			return errors.New("accepts 1 arg(s)")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(File)
		var filename string
		var err error
		var argument string

		if File != "" {
			argument = File
		} else {
			argument = args[0]
		}

		fileExists, err := util.FileExists(argument)
		if err != nil {
			fmt.Println(err)
		}
		if fileExists {
			filename, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())

			}
		} else {
			fmt.Printf("File %v doest not Exists", argument)
			return
		}

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		util.Unzip(filename, wd)

		os.Chdir(util.FilenameWithoutExtension(filename))

		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		commandCode := exec.Command("code", wd)
		err = commandCode.Run()

		if err != nil {
			log.Fatal("VS Code executable file not found in %PATH%")
		}
	},
}

// func init() {
// 	rootCmd.PersistentFlags().StringVarP(&dest, "dest", "d", "", "Pass a destination where you want to copy")

// 	rootCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "Pass a Source from where you want to copy")
// }

func main() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}

}
