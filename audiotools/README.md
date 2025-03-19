# Audio tools

## Example

```go
package main

import (
  "fmt"
  "log"

  "github.com/jtclarkjr/sysutils/audiotools"
)

func main() {
  videoPath := "sample_video.mp4" // Replace with your video file path
  outputDir := "output_audio"     // Replace with your desired output directory

  wavPath, err := audio.ExtractAudio(videoPath, outputDir)
  if err != nil {
    log.Fatalf("Error extracting audio: %v", err)
  }

  fmt.Printf("Audio extracted to: %s\n", wavPath)
}
```
