package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/jmoiron/jsonq"
	"github.com/tbruyelle/hipchat-go/hipchat"
)

// Check error
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Exec shell command
func Run(command string) {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("command failed: %s\n", command)
		fmt.Println(err)
		os.Exit(1)
	}
}

// Exec shell command muted
func RunMuted(command string) {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("command failed: %s\n", command)
		fmt.Println(err)
		os.Exit(1)
	}
}

// Check if a file exist
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// Notify Hipchat room (value from config.json)
func HipchatNotify(message string) bool {
	room := getConfigKey("hipchatRoom")
	token := getConfigKey("hipchatToken")
	c := hipchat.NewClient(token)
	notifRq := &hipchat.NotificationRequest{Message: message}
	_, err := c.Room.Notification(room, notifRq)
	if err != nil {
		panic(err)
	}
	return true
}

func getConfigKey(configKey string) string {
	ConfigPath := "config/env-config.json"
	b, err := ioutil.ReadFile(ConfigPath) // just pass the file name
	Check(err)
	str := string(b) // convert content to a 'string'
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(str))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	brutJson, err := jq.String("config", configKey)
	configKeyStr := string(brutJson)
	return configKeyStr
}
