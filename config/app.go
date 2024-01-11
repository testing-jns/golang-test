package config

import (
	"os"
	"log"
)

var PORT string = "8080"

var BASE_PATH = func() string {
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	
	return path
}

func SetPort(port string) {
    PORT = port
}