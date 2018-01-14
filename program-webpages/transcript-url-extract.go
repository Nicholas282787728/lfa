package main

import (
	"bufio"
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

		content, _ := ioutil.ReadFile(f.Name())
		url := ExtractTranscriptUrl(string(content))

		if url != "" {
			fmt.Printf("wget -O %s %s\n", f.Name(), url) // crufty but effective; a list of wget commands to get the page and save it as the PID
		} else {
			fmt.Println("")
		}

	}
}

func AcceptPipedStrings() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		url := ExtractTranscriptUrl(scanner.Text())
		if url != "" {
			fmt.Println(url)
		}
	}

}

func ExtractTranscriptUrl(input string) string {

	start := `promotion__title"> <a href="`
	end := `" class="br-blocklink__link promotion__link" data-linktrack="promo_title">`

	a := strings.Index(input, start)
	b := strings.Index(input, end)

	if a > 0 && b > 0 {
		return input[a+len(start) : b]
	}

	return ""
}
