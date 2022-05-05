package gorude

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func (target *HttpHeader) ParseFile(filename string) *HttpHeader {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// For grabbing first line
	first := true
	// For grabbing first line

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		// For grabbing first line
		if first {
			requestMethod := strings.Split(text, " ")
			target.Type = requestMethod[0]
			target.Endpoint = requestMethod[1]
			target.HttpVersion = requestMethod[2]
			first = false
			continue
		}
		// For grabbing first line

		if text == "" {
			continue
		}
		// Header set section
		if strings.Contains(text, ":") {
			parsedHeaders := strings.SplitN(text, ":", 2)
			target.Headers[parsedHeaders[0]] = parsedHeaders[1]
		} else {
			target.PostData += text
		}
	}

	target.Domain = strings.TrimSpace(target.Headers["Host"])

	return target
}

func LoadList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
