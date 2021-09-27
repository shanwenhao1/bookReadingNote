package code__test

import (
	"bookReadingNote/aGoExpertProgramming/chapter5/code"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestChannelExample(t *testing.T) {
	code.ChannelExample()
}

func TestChannelTimeOut(t *testing.T) {
	a := assert.New(t)

	status := code.ChannelTimeOut(2*time.Second, 3*time.Second)
	a.Equal(0, status, "mock channel exec normal failed")

	status = code.ChannelTimeOut(3*time.Second, 2*time.Second)
	a.Equal(-1, status, "mock channel exec timeout failed")
}
