package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/appconf/appconf"

	//注册storage驱动
	_ "github.com/appconf/storage/redis"
)

func cfgParser(filename string) (cfg appconf.Config, err error) {
	_, err = toml.DecodeFile(filename, &cfg)
	return
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func main() {
	var filename = flag.String("cfg", "config.toml", "the configuration file")
	flag.Parse()

	cfg, err := cfgParser(*filename)
	if err != nil {
		exit(err)
	}

	if err = appconf.Run(cfg); err != nil {
		exit(err)
	}

	os.Exit(0)
}
