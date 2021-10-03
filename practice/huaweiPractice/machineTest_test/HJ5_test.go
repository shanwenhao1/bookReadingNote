package machineTest

import (
	"bookReadingNote/practice/huaweiPractice/machineTest"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrToInt10(t *testing.T) {
	a := assert.New(t)

	testS := []string{"0xA", "0xAA"}
	exceptS := []string{"10", "170"}
	for pos, str := range testS {
		result, err := machineTest.StrToInt10(str)
		if err != nil {
			t.Errorf(fmt.Sprintf("%s, %s", "test failed", err.Error()))
		}
		a.Equal(exceptS[pos], result, "test failed, StrToInt10 error")
	}
}
