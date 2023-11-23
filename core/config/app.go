package config

import (
	"os"
	"log"
)

const PORT string = "8080"

var BASE_PATH = func() string {
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	
	return path
}

