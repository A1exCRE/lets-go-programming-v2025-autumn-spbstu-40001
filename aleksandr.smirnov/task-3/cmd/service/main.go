package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/A1exCRE/task-3/internal/config"
	"github.com/A1exCRE/task-3/internal/jsonsaver"
	models "github.com/A1exCRE/task-3/internal/valcurs"
	"github.com/A1exCRE/task-3/internal/xmlloader"
)

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Printf("Error: %v\n", rec)
			os.Exit(1)
		}
	}()

	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if configPath == "" {
		panic("config file is required: use --config flag")
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf("load config: %v", err))
	}

	var valCurs models.ValCurs
	if err := xmlloader.LoadFile(cfg.InputFile, &valCurs); err != nil {
		panic(fmt.Sprintf("load XML data: %v", err))
	}

	valCurs.SortByValueDesc()

	if err := jsonsaver.SaveFile(valCurs.Valutes, cfg.OutputFile); err != nil {
		panic(fmt.Sprintf("save JSON result: %v", err))
	}
}
