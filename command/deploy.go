package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var ()

func CmdDeploy(c *cli.Context) {
	// Check config json
	configPath := "config/env-config.json"
	checkConfigJson := Exists(configPath)
	if checkConfigJson == false {
		fmt.Println("The config.json file is not present in the current folder")
		os.Exit(1)
	}
	// Check docker-compose
	//v := "docker-compose version"
	//RunMuted(v)
	// Check docker run
	x := "docker version"
	RunMuted(x)

	// Check args and set variables (localhost/dev/integration/master)

	// Set all config from json

	// Deploy stack
}
