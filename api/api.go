package api

import (
    "fmt"
    "net/http"
    "encoding/json"
    "math/rand"
    "time"
)

type Fortune struct {
    Name     string `json:"Name"`
    Describe string `json:"Describe"`
}

type List struct {
    ListName string    `json:"ListName"`
    Data     []Fortune `json:"Data"`
}

func Setup() List{
    daikiti := Fortune{
        Name: "大吉",
        Describe: "おめでとう！願い事がなんでも叶うよ！",
    }
    tyukiti := Fortune{
        Name: "中吉",
        Describe: "いいことありそう！",
    }
    shokiti := Fortune{
        Name: "小吉",
        Describe: "たぶんいいことあるよ！",
    }
    kyo := Fortune{
        Name: "凶",
        Describe: "強く生きて",
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
    return list
}

func GetList(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    list := Setup()

    json, err := json.Marshal(list)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintln(w, string(json))
}

func GetRandom(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    list := Setup()
    rand.Seed(time.Now().UnixNano())

    json, err := json.Marshal(list.Data[rand.Intn(len(list.Data))])
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintln(w, string(json))
}
