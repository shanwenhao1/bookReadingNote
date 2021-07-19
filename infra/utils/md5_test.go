package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateMd5(t *testing.T) {
	a := assert.New(t)

	data := "appSecretTest0x7"
	md5Str := GenerateMd5(data)
	fmt.Println(md5Str)
	a.Equal("ebe0f75403e429d3279e99ffb8af9d25", md5Str, "generate md5 str error")
}
