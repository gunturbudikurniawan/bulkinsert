package config

import (
	"io/ioutil"
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type Global struct {
	App Application
}

type Application struct {
	TLoc *time.Location
}

var (
	Glb Global
	App Application
)

// Load config from file
func Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("Error: Config file failed to read file -", err)
		return err
	}

	err = yaml.Unmarshal(data, &Glb)
	if err != nil {
		log.Println("Error: Config file failed to Unmarshal -", err)
		return err
	}

	App = Glb.App

	App.TLoc, _ = time.LoadLocation(AppTimeZone)

	return nil
}
