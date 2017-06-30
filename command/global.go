package command

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
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

func GetConfigPath() string {
	return "config/env-config.json"
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
	room := GetConfigKey("hipchatRoom")
	token := GetConfigKey("hipchatToken")
	c := hipchat.NewClient(token)
	notifRq := &hipchat.NotificationRequest{Message: message}
	_, err := c.Room.Notification(room, notifRq)
	if err != nil {
		panic(err)
	}
	return true
}

func GetConfigKey(configKey string) string {
	configPath := GetConfigPath()
	b, err := ioutil.ReadFile(configPath) // just pass the file name
	Check(err)
	str := string(b) // convert content to a 'string'
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(str))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	brutJson, err := jq.String("config", configKey)
	fmt.Printf(brutJson)
	configKeyStr := string(brutJson)
	return configKeyStr
}

func ParseTemplate(from string, to string) {
	t, err := template.ParseFiles(from)
	Check(err)
	f, err := os.Create(to)
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	// Helm template config
	config := map[string]string{
		"siteURL":       "",
		"siteMD5":       "",
		"rootPassword":  "",
		"volumePathDB":  "",
		"volumePathWeb": "",
	}
	err = t.Execute(f, config)
	Check(err)
	f.Close()
}
