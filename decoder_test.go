package hdrx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoder(t *testing.T) {
	type test struct {
		name    string
		input   []string
		decoded Headers
		err     error
	}
	tests := []test{
		{
			"empty",
			[]string{},
			Headers{},
			nil,
		},
		{
			"empty blank",
			[]string{
				"",
			},
			Headers{},
			nil,
		},
		{
			"non-empty blank",
			[]string{
				"    ",
			},
			Headers{},
			nil,
		},
		{
			"single, eof terminated",
			[]string{
				"key: value",
			},
			Headers{
				{"key", "value"},
			},
			nil,
		},
		{
			"single, LF terminated",
			[]string{
				"key: value",
				"",
			},
			Headers{
				{"key", "value"},
			},
			nil,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			input := strings.Join(test.input, "\n")
			r := strings.NewReader(input)
			dec := NewDecoder(r)
			headers, err := dec.ReadHeaders()
			if test.err != nil {
				panic("todo: error tests")
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assert.Equal(t, test.decoded, headers)
		})
	}
}
