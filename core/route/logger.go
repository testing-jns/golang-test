package route

import (
    "log"
    "app/core/terminal/color"
)

func Log(method string, path string, code uint) {
    var color_route_log string = color.Green
    
    if code != 200 {
        color_route_log = color.Red
    }
    
    log.Printf("%s%s - %d - %s %s", color_route_log, method, code, path, color.Reset)
}