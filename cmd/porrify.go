package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/porrify/porrify_api"
)

var goVersion string = runtime.Version()
var version string = ""
var buildDate string = ""

func main() {

	versionFlag := flag.Bool("v", false, "prints current version")
	flag.Parse()
	if *versionFlag {
		fmt.Println("go version ", goVersion)
		fmt.Println("porrify version ", version)
		fmt.Println("build date ", buildDate)
		os.Exit(0)
	}

	config := new(porrify.Config)
	readConfig(config)

	porrify.Run(config)
}

func readConfig(config *porrify.Config) {
	//Read env variables

	// Databases
	environmentVar := os.Getenv("PORRIFY_MYSQL_USER")
	if environmentVar == "" {
		fmt.Println("PORRIFY_MYSQL_USER not set and is required")
		os.Exit(1)
	}
	config.MysqlUser = environmentVar

	environmentVar = os.Getenv("PORRIFY_MYSQL_PASSWORD")
	if environmentVar == "" {
		fmt.Println("PORRIFY_MYSQL_PASSWORD not set and is required")
		os.Exit(1)
	}
	config.MysqlPassword = environmentVar

	environmentVar = os.Getenv("PORRIFY_MYSQL_HOST")
	if environmentVar == "" {
		config.MysqlHost = "localhost"
	} else {
		config.MysqlHost = environmentVar
	}

	environmentVar = os.Getenv("PORRIFY_MYSQL_PORT")
	if environmentVar == "" {
		config.MysqlPort = "3306"
	} else {
		config.MysqlPort = environmentVar
	}

	environmentVar = os.Getenv("PORRIFY_MYSQL_DB")
	if environmentVar == "" {
		fmt.Println("PORRIFY_MYSQL_DB not set and is required")
		os.Exit(1)
	}
	config.MysqlDB = environmentVar
}
