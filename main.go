package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	rundeck "github.com/lusis/go-rundeck/src/rundeck.v17"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	projectid = kingpin.Arg("projectid", "").Required().String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	data, err := client.ListJobs(*projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("%s\n", err)
		} else {
			var prettyJSON bytes.Buffer
			json.Indent(&prettyJSON, jsonData, "", "  ")
			println(string(prettyJSON.Bytes()))
		}
	}
}
