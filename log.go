package tcolorlog

import "fmt"

func PrintError(message string) {
	fmt.Println(Red + "ERROR" + Reset + " " + message)
}