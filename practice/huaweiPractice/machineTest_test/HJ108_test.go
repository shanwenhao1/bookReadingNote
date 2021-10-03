package machineTest

import (
	"bookReadingNote/practice/huaweiPractice/machineTest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLeastComMul(t *testing.T) {
	a := assert.New(t)

	testI := [][]int{{5, 7}}
	exceptI := []int{35}

	for pos, num := range testI {
		res := machineTest.LeastComMul(num[0], num[1])

		a.Equal(exceptI[pos], res, "test failed, LeastComMul error")
	}
}
