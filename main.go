package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func output(data interface{}, template string) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("COGCMD_ERROR: %s\n", err)
	} else {
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, jsonData, "", "  ")
		if len(template) > 0 {
			fmt.Printf("COGCMD_INFO: Using template %s\n", template)
			fmt.Printf("COG_TEMPLATE: %s\n", template)
		}
		println("JSON")
		println(string(prettyJSON.Bytes()))
	}
}

func listJobs() {
	projectid := getProject()
	fmt.Printf("COGCMD_INFO: Listing jobs in project %s\n", projectid)

	client := rundeck.NewClientFromEnv()
	data, err := client.ListJobs(projectid)
	if err != nil {
		fmt.Printf("COGCMD_ERROR: %s\n", err)
		fmt.Printf("ERROR: Cannot list jobs in project %s\n", projectid)
	} else {
		output(&data, "joblist")
	}
}

func runJob(jobArray []string, args string) {
	projectid := getProject()
	jobname := strings.Join(jobArray, " ")
	fmt.Printf("COGCMD_INFO: Running %s in project %s\n", jobname, projectid)

	client := rundeck.NewClientFromEnv()
	data, err := client.FindJobByName(jobname, projectid)
	if err != nil {
		fmt.Printf("COGCMD_ERROR: %s\n", err)
		fmt.Printf("ERROR: Cannot find job %s in project %s\n", jobname, projectid)
	} else {
		fmt.Printf("COGCMD_INFO: Found job %s\n", data.ID)
		jobopts := rundeck.RunOptions{
			RunAs:     "",
			LogLevel:  "",
			Filter:    "",
			Arguments: args,
		}
		fmt.Printf("COGCMD_INFO: Running job %s with args %s\n", data.ID, args)
		res, err := client.RunJob(data.ID, jobopts)
		if err != nil {
			fmt.Printf("COGCMD_ERROR: %s\n", err)
			println("ERROR: Cannot run job %s in project %s", data.ID, projectid)
			os.Exit(1)
		} else {
			executionURL := fmt.Sprintf("%s/project/%s/execution/show/%s",
				os.Getenv("RUNDECK_URL"), projectid, res.Executions[0].ID)
			fmt.Printf("Job %s is %s: %s\n",
				res.Executions[0].ID, res.Executions[0].Status, executionURL)
		}
	}
}

func main() {
	flag.Parse()
	command := flag.Args()[0]
	args := getArgs()
	if len(command) > 0 {
		fmt.Printf("COGCMD_INFO: Running command: %s\n", command)
		switch command {
		case "list-jobs":
			listJobs()
		case "run-job":
			jobArgs := os.Getenv("COG_OPT_ARGS")
			runJob(args, jobArgs)
		default:
			fmt.Printf("COGCMD_ERROR: unknown command: %s\n", command)
			fmt.Printf("ERROR: unknown command: %s\n", command)
			os.Exit(1)
		}
	} else {
		println("COGCMD_ERROR: missing required arguments.")
		println("ERROR: missing required arguments.")
		os.Exit(1)
	}
}
