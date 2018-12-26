package cmd

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-toschool/rom/templates"
	"github.com/spf13/cobra"
)

func insert(s []string, at int, val string) []string {
	// Make sure there is enough room
	s = append(s, "")
	// Move all elements of s up one slot
	copy(s[at+1:], s[at:])
	// Insert the new element at the now free position
	s[at] = val
	return s
}

func createModule(module string, actions []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// check folder
	folderPath := pwd + "/" + module
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0755)
		if err != nil {
			return err
		}
	}

	/* STATE file */
	path := folderPath + "/" + module + ".state.js"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t := templates.StateBase()
		err = ioutil.WriteFile(path, []byte(t), 0777)
		if err != nil {
			return err
		}
	}

	// read file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// scanner on file
	var lines []string
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		var template string
		if strings.Contains(scanner.Text(), "// --end-actions-types--") {
			template = templates.StateActionTypes(actions)
		} else if strings.Contains(scanner.Text(), "// --end-actions-creators--") {
			template = templates.StateActionCreators(actions)
		} else if strings.Contains(scanner.Text(), "// --end-initial-state--") {
			template = templates.StateInitialState(actions)
		} else if strings.Contains(scanner.Text(), "// --end-reducers--") {
			template = templates.StateEndReducers(actions)
		}

		templateLines := strings.Split(template, "\n")
		for _, value := range templateLines {
			lines = insert(lines, i, value)
			i++
		}

		lines = append(lines, scanner.Text())
		i++
	}
	if scanner.Err() != nil {
		return err
	}

	// write file
	out := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(out), 0777)
	if err != nil {
		return err
	}

	/* EPIC file */
	path = folderPath + "/" + module + ".epic.js"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t := templates.StateBase()
		err = ioutil.WriteFile(path, []byte(t), 0777)
		if err != nil {
			return err
		}
	}

	// read file
	file, err = os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// scanner on file
	// var lines []string
	scanner = bufio.NewScanner(file)
	i = 0
	for scanner.Scan() {
		var template string
		template = templates.Epic()
		// if strings.Contains(scanner.Text(), "// --end-actions-types--") {
		// } else if strings.Contains(scanner.Text(), "// --end-actions-creators--") {
		// 	template = templates.StateActionCreators(actions)
		// } else if strings.Contains(scanner.Text(), "// --end-initial-state--") {
		// 	template = templates.StateInitialState(actions)
		// } else if strings.Contains(scanner.Text(), "// --end-reducers--") {
		// 	template = templates.StateEndReducers(actions)
		// }

		templateLines := strings.Split(template, "\n")
		for _, value := range templateLines {
			lines = insert(lines, i, value)
			i++
		}

		lines = append(lines, scanner.Text())
		i++
	}
	if scanner.Err() != nil {
		return err
	}

	// write file
	out = strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(out), 0777)
	if err != nil {
		return err
	}

	//TODO(ca): implements container file
	//TODO(ca): implements component file
	return nil
}

//TODO(mt): change implementation, rom create --module <module_name> <action_1>,<action_2>,…,<action_n>
func parseCreateFlag(opt []string) (string, []string) {
	moduleName := opt[0]
	moduleExtensions := opt[1:len(opt)]

	return moduleName, moduleExtensions
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create --module",
	Short: "example: rom create --module <module_name>,<action_1>,<action_2>,…,<action_n>",
	Run: func(cmd *cobra.Command, args []string) {
		moduleOption, _ := cmd.Flags().GetStringSlice("module")
		mName, mOpt := parseCreateFlag(moduleOption)
		createModule(mName, mOpt)
	},
}

func init() {
	createCmd.Flags().StringSlice("module", []string{"example-module", "foo", "bar"}, "hola mundo")
	rootCmd.AddCommand(createCmd)
}
