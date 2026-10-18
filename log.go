package tcolorlog

import (
	"fmt"
	"io"
	"os"
)

func logf(w io.Writer, color, label, format string, args ...any) {
	fmt.Fprintf(w, color + " " + label + " " + reset + " " + format + "\n", args...)
}

func PrintError(format string, args ...any) {
	logf(os.Stderr, red, "ERROR", format, args...)
}

func PrintSuccess(format string, args ...any) {
	logf(os.Stdout, green, "SUCCESS", format, args...)
}

func PrintWarn(format string, args ...any) {
	logf(os.Stdout, yellow, "WARN", format, args...)
}

func PrintInfo(format string, args ...any) {
	logf(os.Stdout, blue, "INFO", format, args...)
}