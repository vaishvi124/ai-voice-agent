package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func TranslateToEnglish(text string, sourceLanguage string) (string, error) {
	apiKey := os.Getenv("SARVAM_API_KEY")

	if apiKey == "" {
		return "", fmt.Errorf("SARVAM_API_KEY not set")
	}

	body := map[string]interface{}{
		"input":                text,
		"source_language_code": sourceLanguage,
		"target_language_code": "en-IN",
		"model":                "mayura:v1",
		"mode":                 "formal",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.sarvam.ai/translate",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("api-subscription-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("translation failed: %s", string(respBytes))
	}

	var result struct {
		TranslatedText string `json:"translated_text"`
	}

	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", err
	}

	if result.TranslatedText == "" {
		return "", fmt.Errorf("translation returned empty text")
	}

	return result.TranslatedText, nil
}
