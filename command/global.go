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

func GetConfigPath() string {
	return "config.json"
}

// Check error
func Check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
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
	cmd.Stderr = os.Stderr
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

func GetJsonKey(key string, value string) string {
	configPath := GetConfigPath()
	b, err := ioutil.ReadFile(configPath) // just pass the file name
	Check(err)
	str := string(b) // convert content to a 'string'
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(str))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	brutJson, err := jq.String(key, value)
	result := string(brutJson)
	return result
}

func flatMap(src map[string]interface{}, baseKey, sep string, dest map[string]string) {
	for key, val := range src {
		if len(baseKey) != 0 {
			key = baseKey + sep + key
		}
		switch val := val.(type) {
		case map[string]interface{}:
			flatMap(val, key, sep, dest)
		case string:
			dest[key] = val
		case fmt.Stringer:
			dest[key] = val.String()
		default:
			//TODO: You may need to handle ARRAY/SLICE
			//simply convert to string using `Sprintf`
			//modify as you needed.
			dest[key] = fmt.Sprintf("%v", val)
		}
	}
}

func ParseJsonAndTemplate(from string, to string) {
	configPath := GetConfigPath()
	b, _ := ioutil.ReadFile(configPath) // just pass the file name
	var m map[string]interface{}
	err := json.Unmarshal([]byte(b), &m)
	if err != nil {
		log.Fatal(err)
	}
	mm := make(map[string]string)
	flatMap(m, "", "_", mm)
	t, err := template.ParseFiles(from)
	Check(err)
	fmt.Println(mm)
	f, err := os.Create(to)
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	err = t.Execute(f, mm)
	Check(err)
	f.Close()
}

// Notify Hipchat room (value from config.json)
func HipchatNotify(message string) {
	room := GetJsonKey("config", "hipchatRoom")
	token := GetJsonKey("config", "hipchatToken")
	fmt.Sprintf(token)
	c := hipchat.NewClient(token)
	notifRq := &hipchat.NotificationRequest{Message: message}
	err, _ := c.Room.Notification(room, notifRq)
	if err != nil {
	}
}
