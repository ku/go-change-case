package changecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	expectations := [][]string{
		[]string{"RaNdOMcasE", "rAnDomCASe"},
		[]string{"Mix iT.dOWN?", "mIX It.Down?"},
		[]string{"PascalCase", "pASCALcASE"},
		[]string{"Iñtërnâtiônàlizætiøn", "iÑTËRNÂTIÔNÀLIZÆTIØN"},
	}
	for _, pair := range expectations {
		s := pair[0]
		expected := pair[1]
		assert.Equal(t, expected, Swap(s))
	}
}
