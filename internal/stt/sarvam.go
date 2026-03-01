package stt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func Transcribe(audioPath string) (string, error) {
	apiKey := os.Getenv("SARVAM_API_KEY")
url := "https://api.sarvam.ai/v1/audio/transcriptions"
	file, err := os.Open(audioPath)
	if err != nil {
		return "", fmt.Errorf("error opening audio file: %v", err)
	}
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	part, err := writer.CreateFormFile("file", "audio.wav")
	if err != nil {
		return "", err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("RAW RESPONSE:", string(body))

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	text, ok := result["text"].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse transcription: %s", string(body))
	}

	return text, nil
}
