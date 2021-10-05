package machineTest

import (
	"bookReadingNote/practice/huaweiPractice/machineTest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrSort(t *testing.T) {
	a := assert.New(t)

	testD := []string{"c d a bb e"}
	exceptD := []string{"a bb c d e"}
	for pos, test := range testD {
		res := machineTest.StrSort(test)

		a.Equal(exceptD[pos], res, "test failed, StrSort error")
	}
}
