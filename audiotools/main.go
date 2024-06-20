package audio

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// ExtractAudio extracts the audio from a video file and saves it as a WAV file.
func ExtractAudio(videoPath, outputDir string) (string, error) {
	if _, err := os.Stat(videoPath); os.IsNotExist(err) {
		return "", fmt.Errorf("video file does not exist: %s", videoPath)
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return "", fmt.Errorf("failed to create output directory: %v", err)
		}
	}

	outputFilePath := filepath.Join(outputDir, filepath.Base(videoPath)+".wav")
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vn", "-acodec", "pcm_s16le", "-ar", "44100", "-ac", "2", outputFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to extract audio: %v", err)
	}

	return outputFilePath, nil
}
