package strlib

import (
	"testing"
	"testify/assert"
	"time"
)

func TestLcase(t *testing.T) {
	actual := Lcase([]string{"ABC", "DEF"})
	expected := "abc def"
	assert.Equal(t, expected, actual)
}

func TestUcase(t *testing.T) {
	actual := Ucase([]string{"abc", "def"})
	expected := "ABC DEF"
	assert.Equal(t, expected, actual)
}

func TestFlipcase(t *testing.T) {
	actual := FlipCase([]string{"abc", "DEF"})
	expected := "ABC def"
	time.Sleep(time.Second*5)
	assert.Equal(t, expected, actual)
}