package xmlFile

import (
	"encoding/xml"
	"io/ioutil"
)

/*
	xml file interface
*/
type XmlFile interface {
	loadFile(path string) ([]byte, error)
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
func (handle XmlHandle) loadFile(path string) ([]byte, error) {
	fileD, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return fileD, nil
}

/*
	parse xml file to XmlFile interface data
*/
func XmlParse(path string, cfgS XmlFile) (XmlFile, error) {
	var (
		fileD []byte
		err   error
	)
	fileD, err = cfgS.loadFile(path)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(fileD, cfgS)
	if err != nil {
		return nil, err
	}
	return cfgS, nil
}
