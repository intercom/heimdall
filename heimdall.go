package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Rule struct {
	Name        string `json:"name"`
	Pattern     string `json:"pattern"`
	Description string `json:"description"`
	Flags       string `json:"flags"`
	Path        string `json:"path"`
}

type Config struct {
	Rules []Rule `json:"rules"`
}

func main() {
	jsonData, fileError := ioutil.ReadFile("heimdall.json")
	if fileError != nil {
		panic(fmt.Sprintf("Failed to read file 'heimdall.json': %v", fileError))
	}

	var config Config
	jsonError := json.Unmarshal(jsonData, &config)
	if jsonError != nil {
		panic(fmt.Sprintf("Failed to parse JSON config: %v", jsonError))
	}

	exitCode := 0
	fmt.Println()
	for _, rule := range config.Rules {
		if RuleHasViolations(rule) {
			exitCode = 1
		}
	}
	if exitCode == 0 {
		fmt.Println("No rules violated!")
	}
	os.Exit(exitCode)
}

func RuleHasViolations(rule Rule) bool {
	fmt.Printf("Checking Rule '%v'\n\n", rule.Name)

	command := exec.Command("ag")
	// supplying flags or path in constructor will cause ag to fail if either is empty
	if len(rule.Flags) > 0 {
		flags := strings.Fields(rule.Flags)
		for _, flag := range flags {
			command.Args = append(command.Args, flag)
		}
	}
	command.Args = append(command.Args, rule.Pattern)
	if len(rule.Path) > 0 {
		command.Args = append(command.Args, rule.Path)
	}
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()

	if err == nil {
		fmt.Printf("\nRule '%v' was violated:\n%v'\n\n", rule.Name, rule.Description)
	}
	return err == nil
}
