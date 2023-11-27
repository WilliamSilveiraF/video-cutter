package person

import (
	"time"
	"encoding/json"
)

type Person struct {
	ID        int		`json:"id"`
	UserID    int       `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    string    `json:"gender"`
	Contact   string    `json:"contact"`
	Birthday  time.Time	`json:"birthday"`
}

func (p *Person) UnmarshalJSON(data []byte) error {
    type Alias Person
    aux := &struct {
        Birthday string `json:"birthday"`
        *Alias
    }{
        Alias: (*Alias)(p),
    }

    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }

    if aux.Birthday != "" {
        parsedTime, err := time.Parse("2006-01-02", aux.Birthday)
        if err != nil {
            return err
        }
        p.Birthday = parsedTime
    }

    return nil
}
