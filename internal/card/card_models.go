package card

type Card struct {
    ID         int    `json:"id"`
    UserID     int    `json:"user_id"`
    Reference  string `json:"reference"`
    Validity   string `json:"validity"`
    CVV        string `json:"cvv"`
    Cardholder string `json:"cardholder"`
}