// Package ascii provides functions for printing ASCII art
// with optional color highlighting.
package ascii

import (
	"strings"
)

// PrintArgs contains parameters for the PrintAscii function.
type PrintArgs struct {
	Str        string
	Characters []string
}

// PrintAscii prints ASCII art based on the given PrintArgs configuration.
func PrintAscii(args *PrintArgs) string {
	var printline strings.Builder
	index := 0

	// Loop through each line of ASCII art (up to 8 lines)
	for index < 8 {

		for _, char := range args.Str {
			character := args.Characters[int(char)-32]
			lines := strings.Split(character, "\n")
			printline.WriteString(lines[index])

		}
		printline.WriteString("\n")
		index++
	}
	return printline.String()
}
