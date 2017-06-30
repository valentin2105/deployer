package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var (
	environmentPassed string
	realEnvironment   string
)

func CmdDeploy(c *cli.Context) {
	// Check config json
	configPath := "config/env-config.json"
	checkConfigJson := Exists(configPath)
	if checkConfigJson == false {
		fmt.Println("The config.json file is not present in the current folder")
		os.Exit(1)
	}
	// Check docker-compose
	v := "docker-compose version"
	RunMuted(v)
	// Check docker run
	x := "docker version"
	RunMuted(x)
	// Check args and set variables (localhost/dev/integration/master)
	environmentPassed := os.Args[2]
	stackPassed := fmt.Sprintf("composes/%s.tmpl", environmentPassed)
	if Exists(stackPassed) == false {
		fmt.Printf("The environment %s doesn't exist in composes/ path. \n", environmentPassed)
	}
	// Set all config from json

	// Deploy stack
}
