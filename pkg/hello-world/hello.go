package helloworld

import (
	log "github.com/sirupsen/logrus"
)

// Hello returns a friendly greeting
func Hello() string {
	log.Info("Hello World")
	return "Hello World"
}
