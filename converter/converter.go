package converter

import (
	"os"
	"os/exec"
)

const CONVERTED_FILES_DIR = "raw_files/"

var cwd string

func getFFmpegArgs(input string, output string) []string {
	return []string{"-y", "-i", input, "-c", "copy", "-acodec", "pcm_s16le", "-f", "s16le", "-ac", "1", "-ar", "65000", output}
}

func Convert(input string) (string, error) {
	if cwd == "" {
		cwd, _ = os.Getwd()
	}

	output := cwd + CONVERTED_FILES_DIR + input
	if _, err := os.Stat(output); err == nil {
		return output, nil
	}

	cmd := exec.Command("ffmpeg", getFFmpegArgs(input, output)...)
	err := cmd.Start()
	if err != nil {
		return "", err
	}

	return output, cmd.Wait()
}
