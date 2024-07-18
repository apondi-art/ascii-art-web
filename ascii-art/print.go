// Package main is the entry point for the application.
// It prints the ascii representation based on arguments passed.
package asciiprint

import (
	"strings"

	"ascii-art-web/ascii-art/ascii"
	"ascii-art-web/ascii-art/banner"
	"ascii-art-web/ascii-art/errs"
)

func CheckPrint(str, filename string) (string, error) {
	var result strings.Builder
	// Check that str & substr are printable characters
	str = strings.ReplaceAll(str, "\r\n", "\n")
	if err := errs.IsPrintableAscii(str); err != nil {
		return "", err
	}

	contentSlice, err := banner.ReadBannerFile(strings.ToLower(filename))
	if err != nil {
		return "", err
	}

	// Split the str & substr by "\\n" to get the string section in each line
	strs := strings.Split(str, "\n")
	count := 0 // tracks empty strings after splitting str with \n

	for _, s := range strs {
		if s == "" {
			count++
			if count < len(strs) {
				result.WriteString("\n")
			}
		}
		args := &ascii.PrintArgs{
			Str:        s,
			Characters: contentSlice,
		}
		result.WriteString(ascii.PrintAscii(args))
	}
	return result.String(), nil
}
