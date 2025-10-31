package main

import (
	"flag"

	"github.com/A1exCRE/task-3/internal/config"
	"github.com/A1exCRE/task-3/internal/jsonsaver"
	"github.com/A1exCRE/task-3/internal/valcurs"
	"github.com/A1exCRE/task-3/internal/xmlloader"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	var valCurs valcurs.ValCurs

	err = xmlloader.LoadFile(cfg.InputFile, &valCurs)
	if err != nil {
		panic(err)
	}

	valCurs.SortByValueDesc()

	err = jsonsaver.SaveFile(valCurs.Valutes, cfg.OutputFile)
	if err != nil {
		panic(err)
	}
}
