package hdrx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoder(t *testing.T) {
	type header struct {
		key, value string
	}
	type test struct {
		name    string
		headers []header
		encoded []string
		err     error
	}
	tests := []test{
		{
			"empty",
			[]header{},
			[]string{},
			nil,
		},
		{
			"simple",
			[]header{
				{"key", "value"},
			},
			[]string{
				"key: value",
			},
			nil,
		},
		{
			"trimming",
			[]header{
				{"key", "   trimmed    "},
			},
			[]string{
				"key: trimmed",
			},
			nil,
		},
		{
			"multi-line",
			[]header{
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
			for _, header := range test.headers {
				err = enc.WriteHeader(header.key, header.value)
			}
			if test.err != nil {
				panic("todo: error tests")
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
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
