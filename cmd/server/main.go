package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"ai-voice-agent/internal/stt"
)

func main() {
	fmt.Println("Starting STT test...")

	audioFolder := "audio"

	files, err := ioutil.ReadDir(audioFolder)
	if err != nil {
		log.Fatal("Error reading folder:", err)
	}

	// Create CSV file to save results
	csvFile, err := os.Create("results.csv")
	if err != nil {
		log.Fatal("Error creating CSV file:", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()
	writer.Write([]string{"File", "Transcription"}) // header

	// Loop through all .wav files (case-insensitive)
	found := false
	for _, f := range files {
		if f.IsDir() {
			continue // skip folders
		}

		ext := strings.ToLower(filepath.Ext(f.Name()))
		if ext != ".wav" {
			continue // skip non-wav files like .DS_Store
		}

		found = true
		audioPath := filepath.Join(audioFolder, f.Name())
		text, err := stt.Transcribe(audioPath)
		if err != nil {
			continue
		}
		fmt.Println(f.Name(), "->", text)
		writer.Write([]string{f.Name(), text})
	}

	if !found {
		fmt.Println("No .wav files found in folder:", audioFolder)
	} else {
		fmt.Println("All transcriptions complete! Results saved to results.csv")
	}
}

