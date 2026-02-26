package stt

import (
	"bytes"
	"os/exec"
	"strings"
)

func Transcribe(filePath string) (string, error) {
	cmd := exec.Command(
		"whisper-cli",
		"-m", "ggml-medium.bin",
		"-f", filePath,
		"-l", "auto",
		"-nt",
		"-np",
		"-ng", // disable GPU to avoid Metal crash
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	result := strings.TrimSpace(out.String())
	return result, nil
}
