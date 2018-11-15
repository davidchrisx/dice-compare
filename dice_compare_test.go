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
}
