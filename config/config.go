package config

import (
	"io/ioutil"
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type Global struct {
	App  Application
	Ex   External `yaml:"external"`
	My   MySQL    `yaml:"mysql"`
	Ra   Rabbit   `yaml:"rabbit"`
	Rd   Redis    `yaml:"redis"`
	Serv Server   `yaml:"server"`
	URL  Url      `yaml:"url"`
}

type Application struct {
	TLoc *time.Location
}

type External struct {
	OSS struct {
		CompanyLogo struct {
			URL string `yaml:"url"`
		} `yaml:"companylogo"`
	} `yaml:"oss"`
}

type MySQL struct {
	Dialect string `yaml:"dialect"`
	DSN     string `yaml:"dsn"`
}
type Rabbit struct {
	Dialect string `yaml:"dialect"`
	DSN     string `yaml:"dsn"`
}

type Redis struct {
	Addr string `yaml:"addr"`
	Db   int    `yaml:"db"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Url struct {
	Scp struct {
		User string `yaml:"user"`
	} `yaml:"scp"`
}

var (
	Glb  Global
	App  Application
	My   MySQL
	Ra   Rabbit
	Rd   Redis
	Serv Server
	URL  Url
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
	My = Glb.My
	Ra = Glb.Ra
	Rd = Glb.Rd
	Serv = Glb.Serv
	URL = Glb.URL

	App.TLoc, _ = time.LoadLocation(AppTimeZone)

	return nil
}
