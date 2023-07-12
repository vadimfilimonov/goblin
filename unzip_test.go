package goblin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnzip(t *testing.T) {
	type testCase struct {
		name  string
		input [][]string
		want  [][]string
	}

	testCases := []testCase{
		{
			name:  "default",
			input: [][]string{{"a", "1", "!"}, {"b", "2", "?"}},
			want:  [][]string{{"a", "b"}, {"1", "2"}, {"!", "?"}},
		},
		{
			name:  "an empty slice",
			input: [][]string{},
			want:  [][]string{},
		},
		{
			name:  "0-tuples",
			input: [][]string{{}, {}},
			want:  [][]string{},
		},
		{
			name:  "1-tuple",
			input: [][]string{{"barney"}, {"fred"}},
			want:  [][]string{{"barney", "fred"}},
		},
		{
			name:  "2-tuples",
			input: [][]string{{"barney", "36"}, {"fred", "40"}},
			want:  [][]string{{"barney", "fred"}, {"36", "40"}},
		},
		{
			name:  "3-tuples",
			input: [][]string{{"barney", "36", "false"}, {"fred", "40", "true"}},
			want:  [][]string{{"barney", "fred"}, {"36", "40"}, {"false", "true"}},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.want, Unzip(tc.input), tc.name)
	}
}
