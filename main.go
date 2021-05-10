package main

import (
	"fmt"
	"srm-dhcp/config"
	"srm-dhcp/instrumentation"
	"srm-dhcp/parser"

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
	go parser.FileLoop()

	http.HandleFunc("/ping", instrumentation.Ping())
	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println(err)
	}

}
