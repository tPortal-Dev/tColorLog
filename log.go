package tcolorlog

import (
	"fmt"
	"io"
	"os"
)

func logf(w io.Writer, color, label, format string, args ...any) {
	fmt.Fprintf(w, color + " " + label + " " + Reset + " " + format, args...)
}

func PrintError(format string, args ...any) {
	logf(os.Stderr, Red, "ERROR", format, args...)
}

func PrintSuccess(format string, args ...any) {
	logf(os.Stdout, Green, "SUCCESS", format, args...)
}

func PrintWarn(format string, args ...any) {
	logf(os.Stdout, Yellow, "WARN", format, args...)
}