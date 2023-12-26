package gopalindrom_test

import (
	"testing"

	gopalindrom "github.com/koropati/go-palindrom"
	"github.com/stretchr/testify/assert"
)

func Test_IsPalindrom(t *testing.T) {
	var resultsPalindrom interface{} = gopalindrom.IsPalindrom("kakak")
	var resultsNotPalindrom interface{} = gopalindrom.IsPalindrom("abang")

	assert.True(t, resultsPalindrom.(bool))
	assert.Equal(t, true, resultsPalindrom)
	assert.Equal(t, false, resultsNotPalindrom)
}
