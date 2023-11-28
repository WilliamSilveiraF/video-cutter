package audio

import (
    "workflow-editor/db"
)

func InsertAudio(a Audio) (int, error) {
    sqlQuery, err := db.ReadSQLFile("internal/audio/sql/insert_audio.sql")
    if err != nil {
        return 0, err
    }

    var audioID int
    err = db.GetDB().QueryRow(sqlQuery, a.UserID, a.Filename, a.Transcription).Scan(&audioID)
    if err != nil {
        return 0, err
    }

    return audioID, nil
}
