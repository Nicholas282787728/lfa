package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Because I couldn't really figure out how to do this in bash
func main() {
	directoryFiles()
}

func directoryFiles() {
	files, _ := ioutil.ReadDir(".")

	for _, f := range files {

		if strings.Contains(f.Name(), ".txt") {

			content, _ := ioutil.ReadFile(f.Name())
			transcript := extractTranscript(string(content))
			transcript2 := extractTranscript2(string(content))
			transcript3 := extractTranscript3(string(content))

			if transcript != "" {

				/*f, _ := os.Create("../transcripts/" + f.Name())
				defer f.Close()

				f.WriteString(transcript)
				*/

			} else if transcript2 != "" {
			} else if transcript3 != "" {

				f, _ := os.Create("../transcripts/" + f.Name())
				defer f.Close()

				f.WriteString(transcript3)

			} else {
				fmt.Println(f.Name())
			}

		}
	}
}

func extractTranscript(input string) string {
	start := "Letters by date"
	end := ". All rights reserved"

	a := strings.Index(input, start)
	b := strings.Index(input, end)

	if a > 0 && b > 0 {
		return strings.TrimSpace(input[a+len(start) : b])
	}

	return ""
}

func extractTranscript2(input string) string {
	start := "Letters by date"
	end := "COMPLETE ACCURACY."

	a := strings.Index(input, start)
	b := strings.Index(input, end)

	if a > 0 && b > 0 {
		return strings.TrimSpace(input[a+len(start):b]) + " " + end
	}

	return ""
}

func extractTranscript3(input string) string {
	start := "Letters by date"
	end := "Listen to the programme"

	a := strings.Index(input, start)
	b := strings.Index(input, end)

	if a > 0 && b > 0 {
		return strings.TrimSpace(input[a+len(start) : b])
	}

	return ""
}
