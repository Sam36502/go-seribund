/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sam36502/go-seribund/backend"
	"github.com/Sam36502/go-seribund/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "seribund <program-file>",
	Short: "Parses and interprets a seribund file",
	Long: `Parses and interprets a seribund file, then
		prints out all registers as ASCII characters
		sorted alphabetically.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print("Input filename required")
			cmd.Help()
			return
		}
		prgData, err := ioutil.ReadFile(args[0])
		if err != nil {
			cmd.Help()
			return
		}

		prog := backend.ParseProgram(string(prgData))
		result := backend.RunProgram(prog, cmd.Flag(config.FL_STEP).Changed)

		if cmd.Flag(config.FL_VALUES).Changed {
			fmt.Print(backend.RegistersValues(result))
		} else {
			fmt.Print(backend.RegistersASCII(result))
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP(config.FL_VALUES, config.FLS_VALUES, false, "Makes the interpreter list out the values of all registers, instead of their ASCII characters.")
	rootCmd.PersistentFlags().BoolP(config.FL_STEP, config.FLS_STEP, false, "Shows the state of memory while running and prompts the user before continuing.")
}
