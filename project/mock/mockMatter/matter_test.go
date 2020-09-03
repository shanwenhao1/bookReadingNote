package mockMatter

import (
	mock_mockMatter "bookReadingNote/project/mock/mockMatter/mock_matter"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGoMockGotFormatter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mif := mock_mockMatter.NewMockIF(ctrl)

	var iface IF = mif

	wantMatcher := gomock.WantFormatter(
		gomock.StringerFunc(func() string { return "blah" }),
		gomock.Eq((6)),
	)
	gotMatcher := gomock.GotFormatterAdapter(
		gomock.GotFormatterFunc(func(i interface{}) string {
			// Leading 0s
			return fmt.Sprintf("%02d", i)
		}),
		wantMatcher,
	)

	//mif.EXPECT().SimpleMethod(gotMatcher)
	mif.EXPECT().VariadicMethod(gotMatcher)

	//iface.SimpleMethod(5)
	iface.VariadicMethod(4, 2)

}
