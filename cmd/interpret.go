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

	"github.com/spf13/cobra"

	"github.com/Sam36502/go-seribund/backend"
	"github.com/Sam36502/go-seribund/config"
)

// interpretCmd represents the interpret command
var interpretCmd = &cobra.Command{
	Use:   "interpret",
	Short: "Parses and interprets a seribund file",
	Long: `Parses and interprets a seribund file with
		various options for debugging along the way.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}
		prgData, err := ioutil.ReadFile(args[0])
		if err != nil {
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

func init() {
	rootCmd.AddCommand(interpretCmd)

	// Here you will define your flags and configuration settings.
	interpretCmd.PersistentFlags().BoolP(config.FL_VALUES, config.FLS_VALUES, false, "Makes the interpreter list out the values of all registers, instead of their ASCII characters.")
	interpretCmd.PersistentFlags().BoolP(config.FL_STEP, config.FLS_STEP, false, "Shows the state of memory while running and prompts the user before continuing.")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// interpretCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// interpretCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
