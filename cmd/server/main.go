package main

import (
	"fmt"

	"ai-voice-agent/internal/audio"
	"ai-voice-agent/internal/llm"
	"ai-voice-agent/internal/stt"
)

func main() {
	fmt.Println("🚀 Starting Voice Agent...")

	audioPath := "audio/mic.wav"

	// 1️⃣ Record audio from microphone
	err := audio.RecordAudio(audioPath)
	if err != nil {
		fmt.Println("❌ Recording error:", err)
		return
	}

	fmt.Println("✅ Recording saved:", audioPath)

	// 2️⃣ Transcribe audio using Sarvam
	text, err := stt.Transcribe(audioPath)
	if err != nil {
		fmt.Println("❌ Error transcribing:", err)
		return
	}

	fmt.Println("\n📝 Full Transcription:", text)

	// 3️⃣ Detect source language
	sourceLanguage := "en-IN"

	for _, r := range text {

		// Gujarati Unicode range
		if r >= '\u0A80' && r <= '\u0AFF' {
			sourceLanguage = "gu-IN"
			break
		}

		// Hindi / Devanagari Unicode range
		if r >= '\u0900' && r <= '\u097F' {
			sourceLanguage = "hi-IN"
			break
		}
	}

	fmt.Println("🔤 Detected language:", sourceLanguage)

	// 4️⃣ Translate to English using Sarvam
	englishText, err := llm.TranslateToEnglish(text, sourceLanguage)
	if err != nil {
		fmt.Println("❌ Translation error:", err)
		return
	}

	fmt.Println("🌍 English:", englishText)

	fmt.Println("\n✅ Voice processing complete!")
}
