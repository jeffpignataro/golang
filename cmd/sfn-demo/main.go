package main

import (
	sfnhandler "golang/pkg/sfn-handler"

	log "github.com/BillHeroInc/logrus"
)

func main() {
	cfg, err := sfnhandler.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	a, err := cfg.GetActivities()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info(a)
}
