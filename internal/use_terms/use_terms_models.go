package use_terms

type UseTerms struct {
    ID          int    `json:"id,omitempty"`
    Version     string `json:"version"`
    Description string `json:"description"`
}

type UseTermsUser struct {
    ID          int    `json:"id,omitempty"`
    Email       string `json:"email"`
    Password    string `json:"password,omitempty"`
    UseTermsID  int    `json:"use_terms_id,omitempty"`
}