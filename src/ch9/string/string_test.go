package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))
	s = "人"
	t.Log(len(s))

	c := []rune(s)
	t.Logf("人 unicode %x", c[0])
	t.Logf("人 UTF8 %x", s)
}

func TestStringRune(t *testing.T) {
	s := "共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d %[1]x", c)
	}
}
