package machineTest

import (
	"bufio"
	"os"
)

func HJStrScan() []string {
	var scanner = bufio.NewScanner(os.Stdin)
	var inputS []string

	// 获取输入
	for {
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			break
		}
		inputS = append(inputS, input)
	}
	return inputS
}
