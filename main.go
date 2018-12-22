package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-toschool/rom/templates"
)

func createModule(module string, actions []string) error {
	t := templates.State(module, actions)

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(pwd+"/"+module+".state.js", []byte(t), 0777)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("rom create --module auth login,signup")

	module := "auth"
	actions := "login,signup,piroka"

	a := strings.Split(actions, ",")
	createModule(module, a)
}
