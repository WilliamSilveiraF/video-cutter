package audio

import (
    "log"
    "workflow-editor/db"
)

func InsertAudio(a Audio) (int, error) {
    sqlQuery, err := db.ReadSQLFile("internal/audio/sql/insert_audio.sql")
    if err != nil {
        return 0, err
    }

    var audioID int
    err = db.GetDB().QueryRow(sqlQuery, a.UserID, a.Filename).Scan(&audioID)
    if err != nil {
        return 0, err
    }

    return audioID, nil
}

func RetrieveAudiosByUserID(userID int) ([]Audio, error) {
    sqlQuery, err := db.ReadSQLFile("internal/audio/sql/retrieve_audios_by_user_id.sql")
    if err != nil {
        return nil, err
    }

    rows, err := db.GetDB().Query(sqlQuery, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var audios []Audio
    for rows.Next() {
        var audio Audio
        if err := rows.Scan(&audio.ID, &audio.UserID, &audio.Filename); err != nil {
            return nil, err
        }
        audios = append(audios, audio)
    }

    return audios, nil
}

func DeleteAudioByID(audioID int) error {
    sqlQuery, err := db.ReadSQLFile("internal/audio/sql/delete_audio_by_id.sql")
    if err != nil {
        return err
    }

    _, err = db.GetDB().Exec(sqlQuery, audioID)
    return err
}

func RetrieveAudioByID(audioID int) (*Audio, error) {
    sqlQuery := "SELECT id, user_id, filename FROM audio WHERE id = $1"
    var audio Audio
    
    log.Println("RetrieveAudioByID 1")
    err := db.GetDB().QueryRow(sqlQuery, audioID).Scan(&audio.ID, &audio.UserID, &audio.Filename)
    if err != nil {
        return nil, err
    }
    log.Println("RetrieveAudioByID 2")
    audio.FilePath = "temp/" + audio.Filename
    log.Println("RetrieveAudioByID 3")
    return &audio, nil
}

