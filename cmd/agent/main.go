package main

import (
	"fmt"
	"log"

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
		log.Fatalf("❌ Recording error: %v", err)
	}

	fmt.Println("✅ Recording saved:", audioPath)

	// 2️⃣ Transcribe audio using Sarvam
	text, err := stt.Transcribe(audioPath)
	if err != nil {
		log.Fatalf("❌ Error transcribing: %v", err)
	}

	fmt.Println("\n📝 Full Transcription:", text)

	// 3️⃣ Detect source language using internal package
	sourceLanguage := stt.DetectLanguage(text)

	fmt.Println("🔤 Detected language:", sourceLanguage)

	// 4️⃣ Translate to English using Sarvam
	englishText, err := llm.TranslateToEnglish(text, sourceLanguage)
	if err != nil {
		log.Fatalf("❌ Translation error: %v", err)
	}

	fmt.Println("🌍 English:", englishText)

	fmt.Println("\n✅ Voice processing complete!")
}
