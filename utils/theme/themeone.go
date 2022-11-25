package theme

import "fmt"

const (
	colorDefault = "\x1b[39m"

	colorRed   = "\x1b[91m"
	colorGreen = "\x1b[32m"
	colorBlue  = "\x1b[94m"
	colorGray  = "\x1b[90m"
)

func Red(s string) string {
	return fmt.Sprintf("%s%s%s", colorRed, s, colorDefault)
}

func Green(s string) string {
	return fmt.Sprintf("%s%s%s", colorGreen, s, colorDefault)
}

func Blue(s string) string {
	return fmt.Sprintf("%s%s%s", colorBlue, s, colorDefault)
}

func Gray(s string) string {
	return fmt.Sprintf("%s%s%s", colorGray, s, colorDefault)
}
