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
	"fmt"
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/mfahrul/arctic/app"
	"github.com/spf13/cobra"
)

var module = app.Modules()

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:   "module [action]",
	Short: "Module operation",
	Long: `Module operation for your project.
	Available action:
	- add
	- list`,
	ValidArgs:             []string{"add", "list"},
	Args:                  matchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "add":
			if project.Projectname == "" {
				app.Amsyong("========== No project found. ==========")
			}
			app.WorkDir = app.Curdir() + string(filepath.Separator)
			m := module.New(project)
			s := structure.New(m)
			m.Model = *s
			project.Modules = append(project.Modules, *m)
			project.ModuleToParse = *m
			m.CopyModule()
			fmt.Println(app.WorkDir)
			project.Save()
			app.BackupConfig()
			project.ParseProject()
			app.RestoreConfig()
			break
		case "list":
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
			} else {
				fmt.Println("======= No module found. =======")
			}
			break
			// case "remove":
			// 	fmt.Println("======= Function not implemented yet. =======")
			// 	break
		}
	},
}

func matchAll(checks ...cobra.PositionalArgs) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		for _, check := range checks {
			if err := check(cmd, args); err != nil {
				return err
			}
		}
		return nil
	}
}

func init() {
	rootCmd.AddCommand(moduleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moduleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moduleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
