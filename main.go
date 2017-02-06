package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	rundeck "github.com/lusis/go-rundeck/src/rundeck.v17"
)

func getProject() string {
	optProject := os.Getenv("COG_OPT_PROJECT")
	if len(optProject) == 0 {
		defProject := os.Getenv("RUNDECK_DEFAULT_PROJECT")
		if len(defProject) == 0 {
			println("ERROR: Project not specified and RUNDECK_DEFAULT_PROJECT not set.")
			os.Exit(1)
		}
		return defProject
	}
	return optProject
}

func getArgs() []string {
	var args []string
	argc, err := strconv.Atoi(os.Getenv("COG_ARGC"))
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		for i := 0; i < argc; i++ {
			argName := "COG_ARGV_" + strconv.Itoa(i)
			args = append(args, os.Getenv(argName))
		}
	}
	return args
}

func listJobs() {
	projectid := getProject()

	client := rundeck.NewClientFromEnv()
	data, err := client.ListJobs(projectid)
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

func main() {
	args := getArgs()
	if len(args) > 0 {
		switch args[0] {
		case "list-jobs":
			listJobs()
		default:
			println("ERROR: unknown command.")
			os.Exit(1)
		}
	} else {
		println("ERROR: missing required arguments.")
		os.Exit(1)
	}
}
