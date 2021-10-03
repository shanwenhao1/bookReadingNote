package machineTest

import (
	"bookReadingNote/practice/huaweiPractice/machineTest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrSplit(t *testing.T) {
	a := assert.New(t)

	testStrS := []string{
		"abc",
		"123456789",
	}
	exceptStr := [][]string{
		[]string{"abc00000"},
		[]string{"12345678", "90000000"},
	}
	for pos, test := range testStrS {
		result := machineTest.StrSplit(test)
		a.Equal(exceptStr[pos], result, "test failed, wrong result")
	}
}
