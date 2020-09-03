package hystrixExample

import (
	"bookReadingNote/project/CircuitAndHystrix/example/hystrixExample/hystrixManager"
	"bookReadingNote/project/CircuitAndHystrix/example/hystrixExample/mock_hystrixManager"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"math/rand"
	//"net/http"
	"testing"
)

/*
	TODO 修复测试
	预计在gomock v1.5.0版本中修复, https://github.com/golang/mock/issues/391
*/
func TestHystrixRun(t *testing.T) {
	// fake run function
	var fakeRun hystrixManager.RunFunc = func() error {
		n := rand.Intn(5)
		if n < 3 {
			return nil
		}
		return errors.New("mock failed")
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHystrixS := mock_hystrixManager.NewMockHystrixI(ctrl)

	// gomock@1.4.3 s存在问题, 详情见: https://github.com/golang/mock/pull/434/files/3bc6cb03d89418b570a5626382612b7bf254f4cf
	mockHystrixS.EXPECT().Run(fakeRun).Return(nil).AnyTimes()

	// begin test
	err := hystrixManager.HystrixRun(mockHystrixS, fakeRun)
	assert.Equal(t, err.Error(), "test error", "fake run failed")
}
