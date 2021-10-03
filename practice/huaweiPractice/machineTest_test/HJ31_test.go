package machineTest

import (
	"bookReadingNote/practice/huaweiPractice/machineTest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitStrReverse(t *testing.T) {
	a := assert.New(t)

	strS := []string{"I am a student", "$bo*y gi!r#l"}
	exceptS := []string{"student a am I", "l r gi y bo"}
	for pos, str := range strS {
		result := machineTest.SplitStrReverse(str)

		a.Equal(exceptS[pos], result, "test failed, SplitStrReverse function error")
	}
}
