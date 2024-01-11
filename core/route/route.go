package route

import (
    "os"
    "net/http"
    "app/config"
)

type routeCallback func() string
type routeLoggerTypes struct {
    method string
    request_path string
    status_code uint
}


func Route(method string, path string, callback routeCallback) {
    var paths[]string = []string{path}
    
    if path != "/" {
        paths = append(paths, path + "/")
    }

    for _, path := range paths {
        routeHandler(method, path, callback)
    }
}

func routeHandler(method string, path string, callback routeCallback) {
    http.HandleFunc(path, func(res http.ResponseWriter, req *http.Request) {

        var url_req_path string = req.URL.Path
        
        var routeLog routeLoggerTypes
        routeLog.method = req.Method
        routeLog.request_path = url_req_path
        routeLog.status_code = http.StatusOK
        
        if url_req_path != path {
            routeLog.status_code = http.StatusNotFound
            uriNotFound(res, routeLog)
            return
        }
        
        if req.Method != method {
            routeLog.status_code = http.StatusMethodNotAllowed
            methodNotAllowed(res, routeLog)
            return
        }
        
        Log(routeLog.method, routeLog.request_path, routeLog.status_code)
        httpResponse(res, callback(), routeLog.status_code)
    })
}

func uriNotFound(res http.ResponseWriter, routeLog routeLoggerTypes) {
    httpResponse(res, View("warnings/404"), routeLog.status_code)
    Log(routeLog.method, routeLog.request_path, routeLog.status_code)
}

func methodNotAllowed(res http.ResponseWriter, routeLog routeLoggerTypes) {
    var response = `{"success": false, "msg": "Method Not Allowed"}`
    httpResponse(res, response, http.StatusMethodNotAllowed)
    Log(routeLog.method, routeLog.request_path, routeLog.status_code)
}


func httpResponse(res http.ResponseWriter, html string, code uint) {
    res.WriteHeader(int(code))
    res.Header().Set("Content-Type", "text/html; charset=utf-8")
    res.Write([]byte(html))
}


// Static File for Favicon Ico
func FaviconHandler() {
    StaticFile("/static/favicon.ico", "/favicon.ico")
}

func StaticFile(static_file string, request_static_path string) {
    http.HandleFunc(request_static_path, func(res http.ResponseWriter, req *http.Request) {
        var file_path = config.BASE_PATH() + static_file
        
        http.ServeFile(res, req, file_path)
        
        var status_code uint = http.StatusOK
        if !isFileExists(file_path) {
            status_code = http.StatusNotFound
        }
        
        Log(req.Method, req.URL.Path, status_code)
    })
}

func isFileExists(file_path string) bool {
    _, err := os.Stat(file_path)
    if err != nil {
        return false
    }
    return true
}

func StaticDirectory(static_directory string, request_static_path string) {
    static_directory = config.BASE_PATH() + static_directory
    fs := http.FileServer(http.Dir(static_directory))
    strip_prefix := http.StripPrefix(request_static_path, fs)
    http.Handle(request_static_path, strip_prefix)
}