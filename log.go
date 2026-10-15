package tcolorlog

import (
	"fmt"
	"os"
)

func PrintError(message string, a ...any) {
	fmt.Fprintf(os.Stderr, Red + " ERROR " + Reset + " " + message, a...)
}

func PrintSuccess(message string, a ...any) {
	fmt.Printf(Green + " SUCCESS " + Reset + " " + message, a...)
}

func PrintWarn(message string, a ...any) {
	fmt.Printf(Yellow + " WARN " + Reset + " " + message, a...)
}