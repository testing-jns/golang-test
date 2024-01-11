package terminal

import (
	_"fmt"
	"flag"
	"app/config"
)

func Cli() {
	port := flag.String("port", config.PORT, "Type your custom port!")
	flag.Parse()
	
	if *port != "" {
		config.SetPort(*port)
	}
}