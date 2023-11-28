package audio

type Audio struct {
    ID            int    `json:"id"`
    UserID        int    `json:"user_id"`
    Filename      string `json:"filename"`
    Transcription string `json:"transcription"`
}