
package main

import (
	"fmt"
	"log"

	"ai-voice-agent/internal/stt"
)

func main() {

	fmt.Println("Starting STT test...")

	text, err := stt.Transcribe("audio.wav")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User said:", text)
}
