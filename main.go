package main

import (
    "log"
    "net/http"

    "github.com/tetymd/lottery/api"
)

func main() {
    log.Println("server start")
    http.HandleFunc("/api/v1/get/list", api.GetList)
    http.HandleFunc("/api/v1/get/random", api.GetRandom)
    http.ListenAndServe(":8000", nil)
}
