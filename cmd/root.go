/*
Copyright © 2022 Mike Kinney

*/
package cmd

import (
	"os"
	"errors"
	"os/exec"
	"fmt"
	"log"
	"runtime"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "meshtastic-flasher",
	Short: "Update LoRa devices bootloader and firmware for Meshtastic",
	Long: `The Meshtastic project https://meshtastic.org enables low cost, low power radios to be used over LoRa for for communications. This utility updates the bootloader and firmware for use with Meshtastic.`,
	Run: func(cmd *cobra.Command, args []string) {

		// create directory for m-flasher
		home, _ := os.UserHomeDir()
		fmt.Println("home:", home)
		slash := "/"
		if runtime.GOOS == "windows" {
			slash = "\\"
		}

		mf := home + slash + "meshtastic-flasher"
		fmt.Println("mf:", mf)
		if _, err := os.Stat(mf); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(mf, os.ModePerm)
			if err != nil {
				fmt.Print(err)
			}
			fmt.Println("Created directory")
		}

		// create python virtual environment, if we need to
		venv := mf + slash + "venv"
		if _, err := os.Stat(venv); errors.Is(err, os.ErrNotExist) {
			if runtime.GOOS == "windows" {
				// command prompt
				cmd := exec.Command("cmd", "/C", "cd", mf, "&", "python", "-m", "venv", "venv")
				if err := cmd.Run(); err != nil {
					log.Fatal("Could not create python virtual environment")
				}
				fmt.Println("Created python virtual environment")
			} else if runtime.GOOS == "darwin" {
				// bash
				cmd := exec.Command("bash", "-c", "cd " + mf + "; python3 -m venv venv")
				if err := cmd.Run(); err != nil {
					log.Fatal("Could not create python virtual environment")
				}
				fmt.Println("Created python virtual environment")
			}

		}

		// run it
		if runtime.GOOS == "windows" {
			// command prompt
			cmd := exec.Command("cmd", "/C", "cd " + mf + "&", "venv\\Scripts\\activate", "&", "python", "-m", "pip", "install", "--upgrade", "pip", "&", "pip", "install", "--upgrade", "meshtastic-flasher", "&", "meshtastic-flasher")
			if err := cmd.Run(); err != nil {
				log.Fatal("Could not run pip commands")
			}
		} else if runtime.GOOS == "darwin" {
			cmd := exec.Command("bash", "-c", "cd " + mf + "; source venv/bin/activate ; python -m pip install --upgrade pip ; pip install --upgrade meshtastic-flasher ; meshtastic-flasher")
			_, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal("Could not run pip commands:", err, "\n")
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
