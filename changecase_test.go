package changecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Equal(t, "TEST_STRING", Constant("testString"))
	assert.Equal(t, "TEST_STRING", Constant("Test String"))
	assert.Equal(t, "TEST_STRING", Constant("Test_String"))
	assert.Equal(t, "TEST_STRING", Constant("Test-String"))
	assert.Equal(t, "A_BETTER_TEST", Constant("a---better__test"))
}

func TestDot(t *testing.T) {
	assert.Equal(t, "test.string", Dot("testString"))
	assert.Equal(t, "test.string", Dot("Test String"))
	assert.Equal(t, "test.string", Dot("Test_String"))
	assert.Equal(t, "test.string", Dot("Test-String"))
	assert.Equal(t, "a.better.test", Dot("a---better__test"))
}

func TestHeader(t *testing.T) {
	assert.Equal(t, "A-Simple-Test", Header("a-simple-test"))
	assert.Equal(t, "This-Is-A-Test", Header("this is a test"))
	assert.Equal(t, "This-Is-A-Test", Header("this_is_a_test"))
	assert.Equal(t, "This-Is-A-Test", Header("this-is-a-test"))
}

func TestIsLower(t *testing.T) {
	assert.Equal(t, IsLower("ALLUPPERCASE"), false)
	assert.Equal(t, IsLower("NotAllLowerCase"), false)
	assert.Equal(t, IsLower("alllowercase"), true)
	assert.Equal(t, IsLower(""), true)
}
func TestIsUpper(t *testing.T) {
	assert.Equal(t, IsUpper("ALLUPPERCASE"), true)
	assert.Equal(t, IsUpper("NotAllUpperCase"), false)
	assert.Equal(t, IsUpper("alllowercase"), false)
	assert.Equal(t, IsUpper(""), true)
}
func TestLower(t *testing.T) {
	assert.Equal(t, "test", Lower("test"))
	assert.Equal(t, "test", Lower("TEST"))
}

func TestLcFirst(t *testing.T) {
	assert.Equal(t, "tEST", LcFirst("TEST"))
	assert.Equal(t, "test", LcFirst("Test"))
}

func TestParam(t *testing.T) {
	assert.Equal(t, "test-string", Param("testString"))
	assert.Equal(t, "test-string", Param("Test String"))
	assert.Equal(t, "test-string", Param("Test_String"))
	assert.Equal(t, "test-string", Param("Test-String"))
	assert.Equal(t, "a-better-test", Param("a---better__test"))
}

func TestPascal(t *testing.T) {
	assert.Equal(t, "TestString", Pascal("testString"))
	assert.Equal(t, "TestString", Pascal("Test String"))
	assert.Equal(t, "TestString", Pascal("Test_String"))
	assert.Equal(t, "TestString", Pascal("Test-String"))
	assert.Equal(t, "FacebookApi", Pascal("Facebook API"))
	assert.Equal(t, "ATestAgain", Pascal("a-test-again"))
	assert.Equal(t, "ABetterTest", Pascal("a---better__test"))
}

func TestPath(t *testing.T) {
	assert.Equal(t, "test/string", Path("testString"))
	assert.Equal(t, "test/string", Path("Test String"))
	assert.Equal(t, "test/string", Path("Test_String"))
	assert.Equal(t, "test/string", Path("Test-String"))
	assert.Equal(t, "a/better/test", Path("a---better__test"))
}

func TestSentence(t *testing.T) {
	assert.Equal(t, "A simple test", Sentence("a-simple-test"))
	assert.Equal(t, "This is a test", Sentence("this is a test"))
	assert.Equal(t, "This is a test", Sentence("this_is_a_test"))
	assert.Equal(t, "This is a test", Sentence("this-is-a-test"))
}
func TestSnake(t *testing.T) {
	assert.Equal(t, "test_string", Snake("testString"))
	assert.Equal(t, "test_string", Snake("Test String"))
	assert.Equal(t, "test_string", Snake("Test_String"))
	assert.Equal(t, "test_string", Snake("Test-String"))
	assert.Equal(t, "a_better_test", Snake("a---better__test"))
}

func TestTitle(t *testing.T) {
	assert.Equal(t, "A Test Sentence", Title("a test sentence"))
	assert.Equal(t, "I Found A Bug", Title("i found a bug"))
	assert.Equal(t, "Some Things", Title(`some things'`))
	assert.Equal(t, "Quotes", Title(`"quotes"`))
	assert.Equal(t, "Hyphen Ness", Title("hyphen-ness"))
}

func TestUpper(t *testing.T) {
	assert.Equal(t, "TEST", Upper("TEST"))
	assert.Equal(t, "TEST", Upper("test"))
}

func TestUcFirst(t *testing.T) {
	assert.Equal(t, "TEST", UcFirst("TEST"))
	assert.Equal(t, "Test", UcFirst("test"))
}
