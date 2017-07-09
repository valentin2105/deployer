package command

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/codegangsta/cli"
)

func CmdDelete(c *cli.Context) {
	// Check config json
	configPath := GetConfigPath()
	checkConfigJson := Exists(configPath)
	if checkConfigJson == false {
		fmt.Println("The config.json file is not present in the current folder")
		os.Exit(1)
	}
	// New Spinner
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	// Check docker-compose
	v := "docker-compose version"
	RunMuted(v)
	// Check docker run
	x := "docker version"
	RunMuted(x)
	// Check args and set variables (localhost/dev/integration/master)
	environmentPassed := os.Args[2]
	stackPassed := fmt.Sprintf("compose/%s.tmpl.yml", environmentPassed)
	if Exists(stackPassed) == false {
		fmt.Printf("The environment compose/%s.tmpl.yml doesn't exist. \n", environmentPassed)
	}
	// delete stack
	parseDest := fmt.Sprintf(".generated/%s.yml", environmentPassed)
	cmdDown := fmt.Sprintf("docker-compose -f %s down", parseDest)
	RunMuted(cmdDown)
	s.Stop()

}
