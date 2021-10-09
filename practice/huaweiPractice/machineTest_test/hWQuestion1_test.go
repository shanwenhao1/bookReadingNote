package machineTest

import (
	"bookReadingNote/practice/huaweiPractice/machineTest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppleDivide(t *testing.T) {
	a := assert.New(t)

	testD := [][]int{{3, 5, 6}, {180, 280, 810}}
	exceptR := []int{11, -1}
	for pos, data := range testD {
		res, err := machineTest.AppleDivide(data)
		if err != nil {
			a.Error(err)
		}
		a.Equal(exceptR[pos], res, "test failed, AppleDivide error")
	}
}
