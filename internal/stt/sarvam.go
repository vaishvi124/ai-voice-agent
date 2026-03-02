package stt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func Transcribe(audioPath string) (string, error) {
	apiKey := os.Getenv("SARVAM_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("SARVAM_API_KEY not set")
	}

	file, err := os.Open(audioPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", audioPath)
	if err != nil {
		return "", err
	}
	io.Copy(part, file)

	// Sarvam may require model/language fields (safe defaults)
	writer.WriteField("model", "saarika:v2.5")
	writer.Close()

	req, err := http.NewRequest(
		"POST",
		"https://api.sarvam.ai/speech-to-text",
		&body,
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("api-subscription-key", apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

client := &http.Client{
    Timeout: 60 * time.Second,
}
resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)

	//fmt.Println("Status Code:", resp.StatusCode)
	//fmt.Println("RAW RESPONSE:", string(respBytes))

	var result map[string]interface{}
	json.Unmarshal(respBytes, &result)

	// Sarvam usually returns: { "transcript": "..." }
	if text, ok := result["transcript"].(string); ok {
		return text, nil
	}

	// fallback if API returns "text"
	if text, ok := result["text"].(string); ok {
		return text, nil
	}

	return "", fmt.Errorf("failed to parse transcription: %s", string(respBytes))
}

