package main

import (
    "app/core/route"
    "app/core/server"
)

func main() {
    route.FaviconHandler()
    route.StaticFile("/static/robots.txt", "/robots.txt")
    route.StaticDirectory("/static/assets/", "/assets/")
    
    route.Route("GET", "/", func() string {
        return route.View("home/index")
    })
    
    route.Route("GET", "/posts", func() string {
        return route.View("posts/index")
    })
    
    route.Route("GET", "/user", func() string {
        return "Ini user"
    })
    
    server.Serve()
}