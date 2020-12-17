package xmlFile

import (
	mockXmlFile "bookReadingNote/infra/tool/file/xmlFile/mock_xmlfile"
	"encoding/xml"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXmlParse(t *testing.T) {
	a := assert.New(t)

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockXmlF := mockXmlFile.NewMockXmlFile(mockCtl)

	type testXml struct {
		XmlHandle
		Foo string `xml:"foo"`
		Bar string `xml:"bar"`
	}

	// mock LoadFile data
	var testXE = new(testXml)
	testXE.Foo = "test foo"
	testXE.Bar = "test bar"
	testB, err := xml.Marshal(testXE)
	if err != nil {
		t.Error("mock test xmlStruct data []byte data failed, err: ", err)
	}
	mockXmlF.EXPECT().LoadFile("").Return(testB, nil)

	// begin test
	var testX = new(testXml)
	err = XmlParse("", mockXmlF, testX)
	if err != nil {
		t.Error("parse xml file data to XmlFile failed, err: ", err)
	}
	a.Equal(testX.Foo, "test foo", "wrong parse data of XmlFile config")
	a.Equal(testX.Bar, "test bar", "wrong parse data of XmlFile config")
}
