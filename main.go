package main

import (
    "fmt"
    "log"
    "net/http"
    _ "encoding/json"
)

func main() {
    log.Println("server start")
    http.HandleFunc("/kuji/get", GetKuji)
    http.ListenAndServe(":8000", nil)
}

func GetKuji(w http.ResponseWriter, r *http.Request) {
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
