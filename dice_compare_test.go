package dicecompare_test

import (
	dc "github.com/davidchrisx/dicecompare"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitBigram(t *testing.T) {
	ma := dc.SplitBigram("night")
	assert.NotNil(t, ma)
}

func TestCompare(t *testing.T) {
	res := dc.Compare("night", "nacht")
	assert.Equal(t, res, 0.2)
	res = dc.Compare("a", "nacht")
	assert.Equal(t, res, 0.0)
	res = dc.Compare("aa", "aa")
	assert.Equal(t, res, 1)
}

func TestCompareList(t *testing.T) {
	res, _ := dc.CompareList("orc lidi", []string{"orc lady", "marc", "high orc", "orc baby"})
	assert.Equal(t, res, "orc lady")
	res, _ = dc.CompareList("orc lidi", []string{"ddd", "ccc", "bbb", "aaa"})
}
