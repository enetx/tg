package ffmpeg

import (
	"bytes"
	"encoding/json"
	"os/exec"

	. "github.com/enetx/g"
)

// VideoMetadata holds structured information about a video file.
type VideoMetadata struct {
	Width    int64
	Height   int64
	Duration String
}

// ffprobeOutput defines the structure for parsing the JSON output from ffprobe.
type ffprobeOutput struct {
	Streams []struct {
		CodecType string `json:"codec_type"`
		Width     int64  `json:"width"`
		Height    int64  `json:"height"`
		Duration  string `json:"duration"` // ffprobe returns duration as a string.
	} `json:"streams"`
}

// GetVideoMetadata extracts video metadata using the ffprobe command-line tool.
// It returns a Result containing either the metadata or an error.
func GetVideoMetadata(videoPath String) Result[*VideoMetadata] {
	cmd := exec.Command("ffprobe",
		"-v", "error", // Only log critical errors from ffprobe.
		"-print_format", "json",
		"-show_streams",
		videoPath.Std(),
	)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return Err[*VideoMetadata](Errorf("failed to run ffprobe: {}, details: {}", err, stderr.String()))
	}

	var ffprobeData ffprobeOutput
	if err := json.Unmarshal(out.Bytes(), &ffprobeData); err != nil {
		return Err[*VideoMetadata](Errorf("failed to parse ffprobe output: {}", err))
	}

	if len(ffprobeData.Streams) == 0 {
		return Err[*VideoMetadata](Errorf("no streams found in file"))
	}

	for _, stream := range ffprobeData.Streams {
		if stream.CodecType == "video" {
			return Ok(
				&VideoMetadata{
					Width:    stream.Width,
					Height:   stream.Height,
					Duration: String(stream.Duration),
				})
		}
	}

	return Err[*VideoMetadata](Errorf("video stream not found"))
}

// GenerateThumbnail creates a thumbnail for a video and returns the path to the new file.
func GenerateThumbnail(videoPath String, seek ...String) Result[*File] {
	thumbPath := videoPath + ".jpg"

	seekTime := "00:00:01.000"
	if len(seek) > 0 && seek[0] != "" {
		seekTime = seek[0].Std()
	}

	cmd := exec.Command("ffmpeg",
		"-i", videoPath.Std(), // Input file.
		"-ss", seekTime, // Use the determined seek time..
		"-vframes", "1", // Capture only a single frame.
		"-y",            // Overwrite output file if it exists.
		thumbPath.Std(), // Output file.
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return Err[*File](Errorf("failed to generate thumbnail with ffmpeg: {}, details: {}", err, stderr.String()))
	}

	return Ok(NewFile(thumbPath))
}
