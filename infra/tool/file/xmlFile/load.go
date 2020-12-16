//go:generate mockgen -destination mock_xmlfile/mock_xmlfile.go bookReadingNote/infra/tool/file/xmlFile XmlFile
package xmlFile

import (
	"encoding/xml"
	"io/ioutil"
)

/*
	xml file interface
*/
type XmlFile interface {
	LoadFile(path string) ([]byte, error)
}

/*
	xml file handle struct
*/
type XmlHandle struct {
}

/*
	load xml file,
		return []byte data if no error,
		else return nil
*/
func (handle XmlHandle) LoadFile(path string) ([]byte, error) {
	fileD, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return fileD, nil
}

/*
	parse xml file to XmlFile interface data
*/
func XmlParse(path string, cfgS XmlFile) error {
	var (
		fileD []byte
		err   error
	)
	fileD, err = cfgS.LoadFile(path)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(fileD, cfgS)
	if err != nil {
		return err
	}
	return nil
}
