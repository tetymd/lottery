package web

import (
    "fmt"
    "time"
    "net/http"
    "math/rand"
    "encoding/json"

    "github.com/tetymd/lottery/api"
)

type Server struct {
    Port string
}

func (s *Server) Start() {
    http.HandleFunc("/api/v1/get", get)
    http.HandleFunc("/api/v1/get/list", get_list)
    http.ListenAndServe(s.Port, nil)
}

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    list := api.Setup()
    rand.Seed(time.Now().UnixNano())

    json, err := json.Marshal(list.Data[rand.Intn(len(list.Data))])
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintln(w, string(json))
}

func get_list(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    list := api.Setup()

    json, err := json.Marshal(list)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintln(w, string(json))
}
