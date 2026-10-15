package tcolorlog

import (
	"fmt"
	"os"
)

func PrintError(message string, a ...any) {
	fmt.Fprintf(os.Stderr, Red + " ERROR " + Reset + " " + message, a...)
}