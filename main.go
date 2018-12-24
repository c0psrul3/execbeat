package main

import (
	execbeat "github.com/c0psrul3/execbeat/beater"
	"github.com/elastic/beats/libbeat/beat"
	"os"
)

var version = "3.3.0"
var name = "execbeat"

func main() {
	err := beat.Run(name, version, execbeat.New)
	if err != nil {
		os.Exit(1)
	}
}
