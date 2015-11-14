package main

import (
	"fmt"
	"github.com/sanbornm/go-selfupdate/selfupdate"
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	version := "1.0"

	updater := &selfupdate.Updater{
		CurrentVersion: version,
		ApiURL:         "http://www.conf.in.th/auto-update/",
		BinURL:         "http://www.conf.in.th/auto-update/",
		DiffURL:        "http://www.conf.in.th/auto-update/",
		Dir:            "update-tmp",
		CmdName:        "go-selfupdate-example", // app name
	}
	if updater != nil {
		fmt.Println("Checking for new versions...")
		go updater.BackgroundRun()
	}
	fmt.Println("Version:", version)
	log.Println("Start file server at :8080")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./"))))
}
