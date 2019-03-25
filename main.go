package main

import (
    "fmt"
    "log"
    "net/http"
    _ "encoding/json"
)

func main() {
    log.Println("server start")
    http.HandleFunc("/api/v1/get/list", GetList)
    http.HandleFunc("/api/v1/get/random", GetRandom)
    http.ListenAndServe(":8000", nil)
}

func GetList(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, `
        [
            { "rarity": "大吉" },
            { "rarity": "中吉" },
            { "rarity": "小吉" },
            { "rarity": "凶" }
        ]
    `)
}

func GetRandom(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, "[{" + "rarity:" + "大吉" + "}]")
}
