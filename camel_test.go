package changecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCamel(t *testing.T) {
	expectations := [][]string{
		[]string{"test", "test"},
		[]string{"TEST", "test"},
		[]string{"test string", "testString"},
		[]string{"Test String", "testString"},
		[]string{"dot.case", "dotCase"},
		[]string{"path/case", "pathCase"},
		// changecase-go does not treat version numbers as an exception.
		[]string{"version 1.2.10", "version1210"},
		[]string{"version 1.21.0", "version1210"},
		[]string{"TestString", "testString"},
		[]string{"simple éxample", "simpleÉxample"},
	}
	for _, pair := range expectations {
		s := pair[0]
		expected := pair[1]
		assert.Equal(t, expected, Camel(s))
	}

	assert.Equal(t, "testString", Camel("*TestString*"))
	assert.Equal(t, "testString", Camel("TestString"))
	assert.Equal(t, "testString", Camel("Test String"))
	assert.Equal(t, "testString", Camel("Test_String"))
	assert.Equal(t, "testString", Camel("Test-String"))
	assert.Equal(t, "facebookApi", Camel("Facebook API"))
	assert.Equal(t, "webkitTransform", Camel("-webkit-transform"))
	assert.Equal(t, "fooBarBaz", Camel("fooBarBaz"))
	assert.Equal(t, "someThings", Camel("some (things)"))
}
