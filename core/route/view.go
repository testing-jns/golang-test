package route

import (
    "os"
    "log"
    "app/config"
)

var root_path string = config.BASE_PATH() + "/views/"

func View(file_path string) string {
    html, err := os.ReadFile(root_path + file_path + ".html")
    
    if err != nil {
        log.Fatal(err)
    }
    
    return string(html)
}