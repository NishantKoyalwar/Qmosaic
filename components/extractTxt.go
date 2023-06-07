package components

import (
	"fmt"
	"regexp"
)

func Remove() {
	str := text

	// Define the regular expression pattern to match the desired text
	re := regexp.MustCompile(`"quote": "([^"]+)"`)

	// Find the submatch using the regular expression pattern
	submatch := re.FindStringSubmatch(str)

	if len(submatch) >= 2 {
		rename = submatch[1]
		fmt.Println(rename)
	} else {
		fmt.Println("Text not found in the given string.")
	}
}
