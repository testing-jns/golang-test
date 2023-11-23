package server

import (
    "fmt"
    "log"
    "net/http"
    "app/core/config"
    "app/core/terminal/color"
)

func Serve() {
    fmt.Printf("%s Starting golang server at port %s%s %s \n", color.White, color.Blue, config.PORT, color.Reset)
    fmt.Println("")
    
    err_mess := http.ListenAndServe(":" + config.PORT, nil)
    
    if err_mess != nil {
        log.Fatal(err_mess)
    }
}