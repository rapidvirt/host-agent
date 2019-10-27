package main

import (
	"github.com/rapidvirt/host-agent/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	cmd.Execute()
}
