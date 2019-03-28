package main

import (
    "os"
    "log"
    "os/signal"

    "github.com/tetymd/lottery/web"
)

func main() {
    log.Println("server start")

    port := "8000"
    s := &web.Server{Port: ":" + port}
    go s.Start()

    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt)
    <-sigChan

    log.Println("server stop")
}
