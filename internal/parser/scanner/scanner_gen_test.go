package scanner

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/parser/scanner/gen"
	"github.com/tomarrell/lbadd/internal/parser/scanner/ruleset"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

func TestGeneratedTestCases(t *testing.T) {
	assert := assert.New(t)

	testfiles, err := filepath.Glob("testdata/*")
	assert.NoError(err)

	for _, testfile := range testfiles {
		tc := gen.From(testfile)

		input := []rune(tc.Input)

		sc := NewRuleBased(input, ruleset.Default)
		index := 0
		for sc.Peek().Type() != token.EOF {
			got := sc.Next()
			want := tc.Output[index]

			assert.Equal(want.Line(), got.Line(), "mismatched attribute 'line' for token %s", got)
			assert.Equal(want.Col(), got.Col(), "mismatched attribute 'col' for token %s", got)
			assert.Equal(want.Offset(), got.Offset(), "mismatched attribute 'offset' for token %s", got)
			assert.Equal(want.Length(), got.Length(), "mismatched attribute 'length' for token %s", got)
			assert.Equal(want.Type(), got.Type(), "mismatched attribute 'type' for token %s", got)
			assert.Equal(want.Value(), got.Value(), "mismatched attribute 'value' for token %s", got)

			index++
		}
	}
}
