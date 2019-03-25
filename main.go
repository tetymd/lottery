package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

type Fortune struct {
    Name     string `json:"Name"`
    Describe string `json:"Describe"`
    Rarity   int    `json:"Rarity"`
    Weight   int    `json:"Weight"`
}

type List struct {
    ListName string    `json:"ListName"`
    Data     []Fortune `json:"Data"`
}

func main() {
    log.Println("server start")
    http.HandleFunc("/api/v1/list", GetList)
    http.HandleFunc("/api/v1/get/random", GetRandom)
    http.ListenAndServe(":8000", nil)
}

func GetList(w http.ResponseWriter, r *http.Request) {
    daikiti := Fortune{
        Name: "大吉",
        Describe: "a",
        Rarity: 3,
        Weight: 10,
    }
    tyukiti := Fortune{
        Name: "中吉",
        Describe: "a",
        Rarity: 2,
        Weight: 30,
    }
    shokiti := Fortune{
        Name: "小吉",
        Describe: "a",
        Rarity: 1,
        Weight: 40,
    }
    kyo := Fortune{
        Name: "凶",
        Describe: "a",
        Rarity: 0,
        Weight: 20,
    }
    list := List{
        ListName: "運勢",
        Data: []Fortune{
            daikiti,
            tyukiti,
            shokiti,
            kyo,
        },
    }

    d, err := json.Marshal(list)
    if err != nil {
        fmt.Println(err)
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, string(d))
}

func GetRandom(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, "[{" + "rarity:" + "大吉" + "}]")
}
