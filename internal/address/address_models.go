package address

type Address struct {
    ID      int    `json:"id"`
    UserID  int    `json:"user_id"`
    Zip     string `json:"zip"`
    Street  string `json:"street"`
    Unit    string `json:"unit"`
    City    string `json:"city"`
    State   string `json:"state"`
}