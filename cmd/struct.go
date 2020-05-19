/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mfahrul/arctic/app"
	"github.com/spf13/cobra"
)

var structure = app.Structs()

// structCmd represents the struct command
var structCmd = &cobra.Command{
	Use:   "struct [action]",
	Short: "Struct operation",
	Long: `Struct operation for your module.
	Available action:
	- add`,
	ValidArgs:             []string{"add"},
	Args:                  matchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "add":
			if project.Projectname == "" {
				app.Amsyong("========== No project found. ==========")
			}
			app.WorkDir = app.Curdir() + string(filepath.Separator)
			var mIndex int = 0
			moduleCount := len(project.Modules)
			if moduleCount > 0 {
				fmt.Println(
					`======================================
Available modules in this project
======================================`,
				)
				for i, m := range project.Modules {
					fmt.Println(i, "|", strcase.ToCamel(m.Name))
				}
				if moduleCount > 1 {
					for {
						scanner := bufio.NewScanner(os.Stdin)
						msg := "Choose module to use [0-" + strconv.Itoa(len(project.Modules)-1) + "] : "
						strInput := strings.ToLower(app.GetInput(msg, *scanner, true))
						mindex, err := strconv.Atoi(strInput)
						if err == nil && (mindex >= 0 && mindex <= len(project.Modules)-1) {
							mIndex = mindex
							break
						}
					}
				}
				module = &project.Modules[mIndex]

			} else {
				fmt.Println("======= No module found. Creating a new one. =======")
				module = module.New(project)
			}
			fmt.Println("Using module:", strcase.ToCamel(module.Name))
			s := structure.New(module)
			if module.Model.Name != "" {
				module.AddStructs = append(module.AddStructs, *s)
			} else {
				module.Model = *s
				project.Modules = append(project.Modules, *module)
			}

			project.ModuleToParse = *module
			module.CopyModule(project)

			project.Save()
			project.ParseProject()
			break
			// case "list":
			// 	fmt.Println("======= Function not implemented yet. =======")
			// 	break
			// case "remove":
			// 	fmt.Println("======= Function not implemented yet. =======")
			// 	break
		}
	},
}

func init() {
	rootCmd.AddCommand(structCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// structCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// structCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
