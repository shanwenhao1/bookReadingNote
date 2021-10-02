package machineTest_test

import (
	"bookReadingNote/practice/huaweiPractice/machineTest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrWordLLen(t *testing.T) {
	a := assert.New(t)

	var testStrSlice = []string{"hello world", "new life"}
	var testResult = []int{5, 4}
	for pos, testStr := range testStrSlice {
		lastLen := machineTest.StrWordLLen(&testStr)
		a.Equal(testResult[pos], lastLen, "test failed, wrong length")
	}
}
