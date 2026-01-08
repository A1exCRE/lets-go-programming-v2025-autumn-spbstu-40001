package main

import (
	"flag"
	"fmt"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	fmt.Printf("Config: %s\n", *configPath)
}
