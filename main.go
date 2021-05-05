package main

import (
	"flag"
	"fmt"
	"os"
)

var cfgPath = flag.String("config", "./config.yaml", "config file path location")

func main() {
	flag.Parse()
	// read config file
	cfg, err := readConfig(*cfgPath)
	if err != nil {
		// error
		os.Exit(1)
	}
	fmt.Println(*cfg)
}
