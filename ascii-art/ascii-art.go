// Package art returns the ascii art representation
// of a string given a banner filename
package art

import (
	"fmt"
	"strings"

	"ascii-art-web/ascii-art/ascii"
	"ascii-art-web/ascii-art/banner"
	"ascii-art-web/ascii-art/errs"
)

// AsciiArt generates the ASCII art representation of the given string `str`
// using the banner specified by `filename`. It returns the generated ASCII
// art as a string and an error if one occurs.
func AsciiArt(str, filename string) (string, error) {
	str = strings.ReplaceAll(str, "\r\n", "\n") // Replace all occurrences of "\r\n" with "\n" to standardize newlines.
	if err := errs.IsPrintableAscii(str); err != nil {
		return "", err
	}

	contentSlice, err := banner.ReadBannerFile(strings.ToLower(filename))
	if err != nil {
		return "", err
	}

	strs := strings.Split(str, "\n")
	count := 0 // tracks empty strings after splitting str with \n
	var art strings.Builder
	for _, s := range strs {
		if s == "" {
			count++
			if count < len(strs) {
				art.WriteString("\n")
				fmt.Println()
			}
		} else {
			args := &ascii.PrintArgs{
				Str:        s,
				Characters: contentSlice,
			}
			art.WriteString(ascii.PrintAscii(args))
		}
	}
	return art.String(), nil
}
