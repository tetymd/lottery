package api

type Ticket struct {
    Name     string `json:"Name"`
    Describe string `json:"Describe"`
}

type TicketList struct {
    ListName string   `json:"ListName"`
    Data     []Ticket `json:"Data"`
}

type Lottery interface {
    // Show()          (TicketList, error)
    Get()           (Ticket, error)
    GetList()       (Ticket, error)
    // Set(TicketList) error
    // Add(Ticket)     error
    // Delete(string)  error
}

func Setup() TicketList{
    daikiti := Ticket{
        Name: "大吉",
        Describe: "おめでとう！願い事がなんでも叶うよ！",
    }
    tyukiti := Ticket{
        Name: "中吉",
        Describe: "いいことありそう！",
    }
    shokiti := Ticket{
        Name: "小吉",
        Describe: "たぶんいいことあるよ！",
    }
    kyo := Ticket{
        Name: "凶",
        Describe: "強く生きて",
    }
    list := TicketList{
        ListName: "運勢",
        Data: []Ticket{
            daikiti,
            tyukiti,
            shokiti,
            kyo,
        },
    }
    return list
}
