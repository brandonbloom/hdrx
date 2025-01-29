package hdrx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoder(t *testing.T) {
	type test struct {
		name    string
		input   Headers
		encoded []string
		err     error
	}
	tests := []test{
		{
			"empty",
			Headers{},
			[]string{},
			nil,
		},
		{
			"simple",
			Headers{
				{"key", "value"},
			},
			[]string{
				"key: value",
			},
			nil,
		},
		{
			"trimming",
			Headers{
				{"key", "   trimmed    "},
			},
			[]string{
				"key: trimmed",
			},
			nil,
		},
		{
			"multi-line",
			Headers{
				{"key", "line 1\nline 2"},
			},
			[]string{
				"key {",
				"  line 1",
				"  line 2",
				"}",
			},
			nil,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			var err error
			var sb strings.Builder
			enc := NewEncoder(&sb)
			for _, header := range test.input {
				err = enc.WriteHeader(header.Key, header.Value)
			}
			if test.err != nil {
				panic("todo: error tests")
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			actual := sb.String()
			expected := strings.Join(test.encoded, "\n")
			if len(expected) > 0 {
				expected += "\n"
			}
			assert.Equal(t, expected, actual)
		})
	}
}
