package gorude

import (
	"regexp"
)

func refactor(line, payload string) string {
	search, _ := regexp.Compile("§.*?§")

	for {
		final := search.FindStringIndex(line)

		if len(final) == 0 {
			break
		}

		before := final[0]
		after := final[1]
		newLine := line[:before] + payload + line[after:]
		line = newLine

	}

	// fmt.Println(line)

	return line
}
