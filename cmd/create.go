// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-toschool/rom/templates"
	"github.com/spf13/cobra"
)

func createModule(module string, actions []string) error {
	t := templates.State(module, actions)

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Print("::::::::")
	fmt.Println(pwd + "/" + module + ".state.js")
	if err := ioutil.WriteFile(pwd+"/"+module+".state.js", []byte(t), 0777); err != nil {
		return err
	}

	return nil
}

func parseCreateFlag(opt []string) (string, []string) {
	moduleName := opt[0]
	moduleExtensions := opt[1:len(opt)]

	return moduleName, moduleExtensions
}

func handleCreateCmd(opt []string) {
	mName, mOpt := parseCreateFlag(opt)
	createModule(mName, mOpt)
	fmt.Println("------")
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "sdfsfsdfsfdsfdsfsf",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("oli")
		moduleOption, _ := cmd.Flags().GetStringSlice("module")
		handleCreateCmd(moduleOption)
		fmt.Println("piroca")
	},
}

func init() {
	createCmd.Flags().StringSlice("module", []string{"example-module", "foo", "bar"}, "hola mundo")
	rootCmd.AddCommand(createCmd)
}
