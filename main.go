package main

import (
	"fmt"
	"srm-ldap/config"
	"srm-ldap/instrumentation"
	"srm-ldap/parser"

	"net/http"
)

const (
	version = "release-0.5"
	name    = "collector"
)

var GitCommit string
var fullVersion = version + "-" + GitCommit

func main() {
	quit := make(chan struct{})
	defer close(quit)
	config.ReadEnv()
	go parser.LdapLoop()

	http.HandleFunc("/ping", instrumentation.Ping())
	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println(err)
	}

}
