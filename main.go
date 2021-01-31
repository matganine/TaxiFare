package main

import (
	"TaxiFare/config"
	"TaxiFare/endpoints"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/spf13/viper"
)

const configFileName = "tf"

func loadConfig() (*config.Configuration, error) {
	v := viper.New()
	config.SetupViper(v, configFileName)
	return config.New(v)
}

func main() {
	flag.Parse() // required for glog flags
	cfg, err := loadConfig()
	if err != nil {
		glog.Fatalf("Configuration could not be loaded: %v", err)
	}
	ridesFilePath := fmt.Sprintf("%s%s", cfg.DataPath, cfg.RidesFile)

	router := httprouter.New()
	ridesEndpoint, err := endpoints.NewRidesEndpoint(ridesFilePath)

	if err != nil {
		glog.Fatalf("Failed to create the rides endpoint handler. %v", err)
	}

	router.GET("/rides", ridesEndpoint)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		glog.Fatalf("Server failed: %v", err)
	}
}
