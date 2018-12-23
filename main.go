//package main

//import (
//"fmt"
//"io/ioutil"
//"os"
//"strings"

//"github.com/go-toschool/rom/templates"
//)

//func createModule(module string, actions []string) error {
//t := templates.State(module, actions)

//pwd, err := os.Getwd()
//if err != nil {
//return err
//}

//err = ioutil.WriteFile(pwd+"/"+module+".state.js", []byte(t), 0777)
//if err != nil {
//return err
//}

//return nil
//}

//func main() {
//fmt.Println("rom create --module auth login,signup")

//module := "auth"
//actions := "login,signup,piroka"

//a := strings.Split(actions, ",")
//createModule(module, a)
//}

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

package main

import "github.com/go-toschool/rom/cmd"

func main() {
	cmd.Execute()
}
