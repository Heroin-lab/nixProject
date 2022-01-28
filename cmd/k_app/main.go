package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/internal/app/appserver"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "config-path", "configs/app_server.toml", "path to config file")
}

func main() {
	flag.Parse()

	servConf := appserver.NewConfig()
	_, err := toml.DecodeFile(confPath, servConf)
	if err != nil {
		logger.Fatal("Config file error", err)
	}

	if err := appserver.Start(servConf); err != nil {
		logger.Fatal("Server starter error", err)
	}

}

//func main() {
//	fmt.Println("Server is listening for http://localhost:8080/api/v1/")
//	http.HandleFunc("/api/v1/", handler)
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Your request is %s!", r.URL.Path[1:])
//	fmt.Printf("Request is %s \n", r.URL.Path[1:]) // logger for console
//}
