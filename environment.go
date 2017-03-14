package environ

import (
	"os"
	"regexp"
)

var environmentVarMatch = regexp.MustCompile(`([\w]+)=(.*)`)

func Map() map[string]string {
	environ := os.Environ()

	result := make(map[string]string, len(environ))

	for _, variable := range environ {
		match := environmentVarMatch.FindAllStringSubmatch(variable, -1)

		if len(match) != 1 {
			continue
		}
		result[match[0][1]] = match[0][2]
	}

	return result
}
