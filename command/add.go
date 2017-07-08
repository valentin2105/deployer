package command

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/codegangsta/cli"
)

var (
	environmentPassed string
	realEnvironment   string
	hipchatMessage    string
)

func CmdAdd(c *cli.Context) {
	// Check config json
	configPath := GetConfigPath()
	checkConfigJson := Exists(configPath)
	if checkConfigJson == false {
		fmt.Println("There is no config.json in the current folder")
		os.Exit(1)
	}
	// New Spinner
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	// Check docker run
	x := "docker version"
	RunMuted(x)
	// Check docker-compose
	v := "docker-compose version"
	RunMuted(v)
	// Check args and set variables (localhost/dev/integration/master)
	environmentPassed := os.Args[2]
	stackPassed := fmt.Sprintf("compose/%s.tmpl.yml", environmentPassed)
	if Exists(stackPassed) == false {
		fmt.Printf("The file compose/%s.tmpl.yml doesn't exist. \n", environmentPassed)
	}
	// Check .generated folder exist
	if _, err := os.Stat(".generated"); os.IsNotExist(err) {
		Run("mkdir .generated")
	}
	// Set all config from json
	parseDest := fmt.Sprintf(".generated/%s.yml", environmentPassed)
	ParseJsonAndTemplate(stackPassed, parseDest)
	// Pull & Deploy stack
	cmdPull := fmt.Sprintf("docker-compose -f %s pull", parseDest)
	RunMuted(cmdPull)
	cmdUp := fmt.Sprintf("docker-compose -f %s up -d", parseDest)
	RunMuted(cmdUp)

	// Run hipothetical Hook
	hookPath := GetJsonKey(environmentPassed, "Hook")
	if hookPath != "" {
		// Wait a little before hook
		Run(hookPath)
	}
	// Notifiy Hipchat of the deployment
	vhost := GetJsonKey(environmentPassed, "Vhost")
	hostname, _ := os.Hostname()
	hipchatMessage := fmt.Sprintf("http://%s just deployed ! (%s)", vhost, hostname)
	HipchatNotify(hipchatMessage)
	s.Stop()
}
