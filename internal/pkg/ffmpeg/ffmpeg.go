package ffmpeg

import (
	"bytes"
	"encoding/json"
	"os/exec"

	"github.com/enetx/g"
)

// VideoMetadata holds structured information about a video file.
type VideoMetadata struct {
	Width    int64
	Height   int64
	Duration g.String
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
// It returns a g.Result containing either the metadata or an error.
func GetVideoMetadata(videoPath g.String) g.Result[*VideoMetadata] {
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
		return g.Err[*VideoMetadata](g.Errorf("failed to run ffprobe: {}, details: {}", err, stderr.String()))
	}

	var ffprobeData ffprobeOutput
	if err := json.Unmarshal(out.Bytes(), &ffprobeData); err != nil {
		return g.Err[*VideoMetadata](g.Errorf("failed to parse ffprobe output: {}", err))
	}

	if len(ffprobeData.Streams) == 0 {
		return g.Err[*VideoMetadata](g.Errorf("no streams found in file"))
	}

	for _, stream := range ffprobeData.Streams {
		if stream.CodecType == "video" {
			return g.Ok(
				&VideoMetadata{
					Width:    stream.Width,
					Height:   stream.Height,
					Duration: g.String(stream.Duration),
				})
		}
	}

	return g.Err[*VideoMetadata](g.Errorf("video stream not found"))
}

// GenerateThumbnail creates a thumbnail for a video and returns the path to the new file.
func GenerateThumbnail(videoPath g.String, seek ...g.String) g.Result[*g.File] {
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
		return g.Err[*g.File](
			g.Errorf("failed to generate thumbnail with ffmpeg: {}, details: {}", err, stderr.String()))
	}

	return g.Ok(g.NewFile(thumbPath))
}
