package audio

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/gordonklaus/portaudio"
)

func RecordAudio(filename string) error {

	const sampleRate = 16000
	const seconds = 5

	portaudio.Initialize()
	defer portaudio.Terminate()

	buffer := make([]int16, 64)

	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(buffer), &buffer)
	if err != nil {
		return err
	}

	stream.Start()
	fmt.Println("🎤 Recording... speak now (5 seconds)")

	var recorded []int16

	for i := 0; i < sampleRate*seconds/len(buffer); i++ {
		err := stream.Read()
		if err != nil {
			return err
		}
		recorded = append(recorded, buffer...)
	}

	stream.Stop()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = writeWAVHeader(file, len(recorded), sampleRate)
	if err != nil {
		return err
	}

	for _, sample := range recorded {
		err := binary.Write(file, binary.LittleEndian, sample)
		if err != nil {
			return err
		}
	}

	fmt.Println("✅ Recording saved:", filename)

	return nil
}

func writeWAVHeader(file *os.File, numSamples int, sampleRate int) error {

	var (
		numChannels   = 1
		bitsPerSample = 16
		byteRate      = sampleRate * numChannels * bitsPerSample / 8
		blockAlign    = numChannels * bitsPerSample / 8
		dataSize      = numSamples * numChannels * bitsPerSample / 8
	)

	file.Write([]byte("RIFF"))
	binary.Write(file, binary.LittleEndian, uint32(36+dataSize))
	file.Write([]byte("WAVE"))

	file.Write([]byte("fmt "))
	binary.Write(file, binary.LittleEndian, uint32(16))
	binary.Write(file, binary.LittleEndian, uint16(1))
	binary.Write(file, binary.LittleEndian, uint16(numChannels))
	binary.Write(file, binary.LittleEndian, uint32(sampleRate))
	binary.Write(file, binary.LittleEndian, uint32(byteRate))
	binary.Write(file, binary.LittleEndian, uint16(blockAlign))
	binary.Write(file, binary.LittleEndian, uint16(bitsPerSample))

	file.Write([]byte("data"))
	binary.Write(file, binary.LittleEndian, uint32(dataSize))

	return nil
}
