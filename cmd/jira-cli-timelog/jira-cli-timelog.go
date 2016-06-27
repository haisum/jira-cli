package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/bgentry/speakeasy"
	"github.com/go-ini/ini"
)

const (
	defaultConf = "jira-cli-config.ini"
)

var (
	config           = flag.String("config", defaultConf, fmt.Sprintf("Config file to use. By default program looks for file %s in current working directory. You can generate a sample config file by executing 'jira-cli-config > %s'", defaultConf, defaultConf))
	requiredConfKeys = []string{"username", "endpoint"}
)

func main() {
	flag.Parse()
	cfg, err := ini.Load(*config)
	if err != nil {
		flag.Usage()
		log.Fatalf("Couldn't load config file %s. %s. Use -config option to supply a valid configuration file.", *config, err)
	}
	for _, key := range requiredConfKeys {
		if !cfg.Section("").HasKey(key) {
			log.Fatalf("One or more of required keys are missing from config %s. Required: %s", *config, strings.Join(requiredConfKeys, ","))
		}
	}
	uKey, _ := cfg.Section("").GetKey("username")
	username := uKey.String()
	epKey, _ := cfg.Section("").GetKey("endpoint")
	endpoint := epKey.String()
	password, err := speakeasy.Ask("Please enter your jira login password: ")
	if err != nil {
		log.Fatalf("Couldn't read password. %s", err)
	}

	log.Printf("username: %s, endpoint: %s, password: %s", username, endpoint, password)
}
