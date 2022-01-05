package gorude

import (
	"regexp"
)

func refactor(line, payload string) string {
	search, _ := regexp.Compile("§.*?§")

	for {
		sonuc := search.FindStringIndex(line)

		if len(sonuc) == 0 {
			break
		}

		before := sonuc[0]
		after := sonuc[1]
		new := line[:before] + payload + line[after:]
		line = new

	}

	// fmt.Println(line)

	return line
}
