package machineTest

import (
	"bookReadingNote/practice/huaweiPractice/machineTest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeDuAndSort(t *testing.T) {
	a := assert.New(t)

	testSlice := []int{2, 2, 1}

	result := machineTest.DeDuAndSort(&testSlice)
	a.Equal([]int{1, 2}, *result, "test failed, result error")
}
