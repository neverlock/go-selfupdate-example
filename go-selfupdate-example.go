package main

import (
	"fmt"
	"github.com/neverlock/utility/selfupdate"
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	version := "1.0"
	updateUrl := "http://www.conf.in.th/auto-update/"
	appName := "go-selfupdate-example"

	go selfupdate.SelfUpdate(version, updateUrl, appName)

	fmt.Println("Version:", version)
	log.Println("Start file server at :8080")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./"))))
}
