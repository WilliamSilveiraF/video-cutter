package audio

import (
    "io/ioutil"
    "context"

    speech "cloud.google.com/go/speech/apiv1"
    speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func createSpeechClient(ctx context.Context) (*speech.Client, error) {
    client, err := speech.NewClient(ctx)
    if err != nil {
        return nil, err
    }
    return client, nil
}

func transcribeAudio(ctx context.Context, client *speech.Client, filePath string) (string, error) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return "", err
    }

    audio := speechpb.RecognitionAudio{
        AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
    }

    config := speechpb.RecognitionConfig{
        Encoding:        speechpb.RecognitionConfig_LINEAR16,
        SampleRateHertz: 16000,
        LanguageCode:    "en-US",
    }

    resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
        Config: &config,
        Audio:  &audio,
    })
    if err != nil {
        return "", err
    }

    if len(resp.Results) == 0 {
        return "", nil
    }
    return resp.Results[0].Alternatives[0].Transcript, nil
}
