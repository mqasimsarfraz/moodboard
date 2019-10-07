package main

import (
	"github.com/MQasimSarfraz/moodboard/pkg/service"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var opts struct {
		Address string `long:"http-address" default:":3080" description:"Address to listen on."`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	service.Run()

}
