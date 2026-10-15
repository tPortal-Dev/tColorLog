package tcolorlog

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPrintError(t *testing.T) {
	oldStderr := os.Stderr

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	os.Stderr = w

	PrintError("%s", "test")

	w.Close()
	os.Stderr = oldStderr

	var output bytes.Buffer
	_, err = io.Copy(&output, r)
	if err != nil {
		t.Fatal(err)
	}
	r.Close()

	expected := Red + " ERROR " + Reset + " " + "test"
	if output.String() != expected {
		t.Errorf("expected result %q, obtained %q", expected, output.String())
	}
}