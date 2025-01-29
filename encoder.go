package hdrx

import (
	"fmt"
	"io"
	"strings"
)

type Encoder struct {
	w     io.Writer
	err   error
	depth int
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (enc *Encoder) WriteHeader(key, value string) error {
	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)

	// TODO: Validate key.

	multiline := false
	if strings.Contains(value, "\n") {
		multiline = true
	}

	enc.print(key)

	if multiline {
		enc.print(" {\n")
		enc.depth += 1
		enc.writeValue(value)
		enc.depth -= 1
		enc.print("}\n")
	} else {
		enc.print(": ")
		enc.writeValue(value)
	}

	return enc.err
}

func (enc *Encoder) print(s string) {
	if enc.err == nil {
		_, enc.err = fmt.Fprint(enc.w, s)
	}
}

func (enc *Encoder) writeValue(s string) error {
	// TODO: escaping.
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		for range enc.depth {
			enc.print("  ")
		}
		enc.print(line)
		enc.print("\n")
	}
	return enc.err
}
