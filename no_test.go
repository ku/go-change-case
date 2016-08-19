package changecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNo(t *testing.T) {
	expectations := [][]string{
		// Single words.
		[]string{"test", "test"},
		[]string{"TEST", "test"},

		// Camel case.
		[]string{"testString", "test string"},
		[]string{"testString123", "test string123"},
		[]string{"testString_1_2_3", "test string 1 2 3"},
		[]string{"x_256", "x 256"},
		[]string{"anHTMLTag", "an html tag"},
		[]string{"ID123String", "id123 string"},
		[]string{"Id123String", "id123 string"},
		[]string{"foo bar123", "foo bar123"},
		[]string{"a1bStar", "a1b star"},

		// Constant case.
		[]string{"CONSTANT_CASE ", "constant case"},
		[]string{"CONST123_FOO", "const123 foo"},

		// Random cases.
		[]string{"FOO_bar", "foo bar"},

		// Non-alphanumeric separators.
		[]string{"dot.case", "dot case"},
		[]string{"path/case", "path case"},
		[]string{"snake_case", "snake case"},
		[]string{"snake_case123", "snake case123"},
		[]string{"snake_case_123", "snake case 123"},

		// Punctuation.
		[]string{`"quotes"`, "quotes"},

		// Space between number parts.
		[]string{"version 0.45.0", "version 0 45 0"},
		[]string{"version 0..78..9", "version 0 78 9"},
		[]string{"version 4_99/4", "version 4 99 4"},

		// Whitespace.
		[]string{"  test  ", "test"},

		// Non-ascii characters.
		[]string{"español", "español"},
		[]string{"Beyoncé Knowles", "beyoncé knowles"},
		[]string{"Iñtërnâtiônàlizætiøn", "iñtërnâtiônàlizætiøn"},

		// Number string input.
		[]string{"something_2014_other", "something 2014 other"},

		// https://github.com/blakeembrey/change-case/issues/21
		[]string{"amazon s3 data", "amazon s3 data"},
		[]string{"foo_13_bar", "foo 13 bar"},

		[]string{"12months", "12months"},
		[]string{"256Colors", "256 colors"},
		[]string{"256colors", "256colors"},
		[]string{"16Colors", "16 colors"},
		[]string{"16colors", "16colors"},
		[]string{"8Colors", "8 colors"},
		[]string{"8colors", "8colors"},
		[]string{"I16n", "i16n"},
		[]string{"facebookAPI", "facebook api"},
	}
	for _, pair := range expectations {
		s := pair[0]
		expected := pair[1]
		assert.Equal(t, expected, No(s))
	}
}
