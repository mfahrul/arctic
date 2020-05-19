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
	"strings"

	"github.com/mfahrul/arctic/app"
	"github.com/spf13/cobra"
)

var project = app.Projects()

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:                   "init",
	Short:                 "Initialize a project",
	Long:                  `Initialize (arctic init) will create a new arctic project.`,
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Example:               "arctic init domain.com/project/name",
	Run: func(cmd *cobra.Command, args []string) {
		project.Projectpath = args[0]
		c := strings.Split(args[0], "/")
		project.Projectname = c[len(c)-1]

		project.New()
		app.WorkDir = app.Curdir() + string(filepath.Separator) + project.Projectname + string(filepath.Separator)

		m := module.New(project)
		s := structure.New(m)
		m.Model = *s
		project.Modules = append(project.Modules, *m)
		project.ModuleToParse = *m
		m.CopyModule(project)

		project.Save()
		// app.WorkDir = app.Curdir() + project.Projectname + string(filepath.Separator)
		project.ParseProject()
		app.MoveFile(app.CfgPath, app.WorkDir+app.CfgPath)
		app.Execute(app.WorkDir+"App", "go", "get", "-u")
		fmt.Println("========== DONE ==========")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
