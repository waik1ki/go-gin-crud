package main

import (
	"flag"
	"go-crud/init/cmd"
)

var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
