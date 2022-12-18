package gopyast

import (
	"io/ioutil"
	"log"
	"strings"
)

type (
	PyParser struct {
		filePath    string
		searchUsage string

		fileContent string
		fileArr     []string
	}

	PyImport     struct{}
	PyFromImport struct{}
	PyClass      struct{}
	PyFunc       struct{}
)

func NewPyParser(filePath, searchUsage string) (pyp *PyParser, err error) {
	pyp = &PyParser{
		filePath: filePath,
	}
	return
}

func (pyp *PyParser) readFile() (err error) {
	var (
		fileContentB []byte
	)
	if fileContentB, err = ioutil.ReadFile(pyp.filePath); err != nil {
		return
	}
	pyp.fileContent = string(fileContentB)
	pyp.fileArr = strings.Split(pyp.fileContent, "\n")
	return
}

func (pyp *PyParser) isImportPresent() (isPresent bool) {
	isPresent = true
	return
}

func (pyp *PyParser) iterateFile() (err error) {
	pyp.startBlockWalk(nil)
	return
}

func (pyp *PyParser) Process() (err error) {
	if err = pyp.readFile(); err != nil {
		return
	}

	if !pyp.isImportPresent() {
		return
	}
	log.Println("Gello", len(pyp.fileArr))
	return
}
