package machineTest

import (
	"bookReadingNote/practice/huaweiPractice/machineTest"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPassCheck(t *testing.T) {
	a := assert.New(t)

	fmt.Println(strings.Count("0000", "000"))
	testS := []string{
		"021Abc9000",
		"021Abc9Abc1",
		"021ABC9000",
		"021$bc9000",
		"021$bc90000",
		"7v0T+6s!(7*)C4RX8*IB85yk+6&~#v6)q$+W3&8-8+",
	}
	exceptR := []bool{
		true,
		false,
		false,
		true,
		false,
		true,
	}
	for pos, testD := range testS {
		fmt.Println(testD)
		result := machineTest.PassCheck(testD)
		fmt.Println(result)
		a.Equal(exceptR[pos], result, "test failed, PassCheck error")
	}
}
