package main

import (
    "log"

    "github.com/tetymd/lottery/web"
)

func main() {
    log.Println("server start")

    port := "8000"

    s := web.Server{Port: ":" + port}
    s.Start()
}
