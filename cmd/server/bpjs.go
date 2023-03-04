package main

import (
	"bpjs/config"
	"bpjs/pkg/handler"
	"bpjs/pkg/myservice"
	"bpjs/pkg/storage/mysql"
	"bpjs/pkg/storage/redis"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {

	goEnv := strings.ToLower(os.Getenv("GO_ENV"))
	if goEnv == "" {
		goEnv = "local"
	}

	// Load config
	LoadConfig(goEnv)

	log.Println(strings.ToUpper(config.AppName), "is warming up ...")

	// Run the server
	run(goEnv)
}

// LoadConfig load config.yaml file from env, user input, or config folder
func LoadConfig(goEnv string) {
	var arg string

	if configFile := os.Getenv("CONFIG_FILE"); configFile != "" {
		arg = configFile
	} else if len(os.Args) == 2 {
		arg = "config/config." + os.Args[1] + ".yaml"
	} else {
		arg = "config/config." + goEnv + ".yaml"
	}

	err := config.Load(arg)
	if err != nil {
		log.Fatal("Error: Config failed to load - ", err)
	}

	log.Println("Load config from", arg)
}

func run(goEnv string) {
	// MySQL setup
	mysql, err := mysql.NewStorage(config.My)
	if err != nil {
		log.Fatal("Error: Database failed to connect (", config.My.DSN, ") - ", err)
	}

	// Redis setup
	redis, err := redis.NewStorage(config.Rd)
	if err != nil {
		log.Fatal("Error: Redis failed to connect (", config.Rd.Addr, ") - ", err)
	}

	// Handler setup
	lister := myservice.NewService(mysql, redis)

	r := handler.Handler(lister)

	host := config.Glb.Serv.Host
	if host == "" {
		host = GetLocalIP()
		config.Glb.Serv.Host = host
	}

	log.Println("Server Running on", goEnv, "environment, (REST APIs) listening on", host+":"+config.Serv.Port)
	log.Fatal("Error: Server failed to run - ", http.ListenAndServe(host+":"+config.Serv.Port, r))
}

func GetLocalIP() string {
	address, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range address {
		// check the address type and if it is not a loopback then display it
		if inet, ok := address.(*net.IPNet); ok && !inet.IP.IsLoopback() {
			if inet.IP.To4() != nil {
				return inet.IP.String()
			}
		}
	}
	return ""
}
