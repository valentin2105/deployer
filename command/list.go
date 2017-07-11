package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func CmdList(c *cli.Context) {
	environmentPassed := os.Args[2]
	stackPath := fmt.Sprintf(".generated/%s.yml", environmentPassed)
	cmdList := fmt.Sprintf("docker-compose -f %s ps", stackPath)
	RunMuted(cmdList)
}
