package machineTest

import (
	"bookReadingNote/practice/huaweiPractice/machineTest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterNum(t *testing.T) {
	a := assert.New(t)

	strSlice := []string{"ABCabc", "ACbCMb"}
	letterSlice := []string{"a", "b"}
	exceptNum := []int{2, 2}
	for pos, str := range strSlice {
		leNum := machineTest.LetterNum(&str, &letterSlice[pos])
		a.Equal(exceptNum[pos], leNum, "test failed, wrong letter count")
	}
}
