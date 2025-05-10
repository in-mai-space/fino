package main

import (
	"fino/internal"
	"flag"
)

func main() {
	cfgPath := parseFlags()

	app, cfg := internal.CreateServer(*cfgPath)

	go func() {
		internal.StartServer(app, cfg)
	}()

	// TODO: handle server shutdown
}

func parseFlags() (configPath *string) {
	configPath = flag.String("config", "", "Specify the path to .env file")
	flag.Parse()
	return
}
