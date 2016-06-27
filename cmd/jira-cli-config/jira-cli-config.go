package main

import "fmt"

const (
	// SampleConfig is sample config file that should be saved in jira-cli-config.ini file in same directory as binary
	SampleConfig = `endpoint=https://jira.myorg.com/jira/rest/api/latest/
username=haisum
`
)

func main() {
	fmt.Println(SampleConfig)
}
