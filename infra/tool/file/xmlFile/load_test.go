package xmlFile

import (
	mockXmlFile "bookReadingNote/infra/tool/file/xmlFile/mock_xmlfile"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestXmlParse(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockXmlF := mockXmlFile.NewMockXmlFile(mockCtl)
	//type testXml struct{
	//	*mockXmlFile.MockXmlFile
	//	TestStr		string		`xml:"test_str"`
	//}
	//var testS = testXml{
	//	mockXmlF,
	//	"test xml",
	//}
	//testB, err := json.Marshal(&testS)
	//if err != nil{
	//	panic(errors.New(fmt.Sprintf("mock test xmlStruct data []byte data failed, err: %v", err)))
	//}
	// mock LoadFile return
	testB, err := json.Marshal(mockXmlF)
	if err != nil {
		panic(errors.New(fmt.Sprintf("mock test xmlStruct data []byte data failed, err: %v", err)))
	}
	mockXmlF.EXPECT().LoadFile("").Return(testB, nil)

	// begin test
	//var resultX = new(testXml)
	err = XmlParse("", mockXmlF)
	if err != nil {
		fmt.Println("----=+++----", err)
		panic(err)
	}
	fmt.Println("-------", mockXmlF)
}
