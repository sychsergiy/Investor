package file

import (
	"io/ioutil"
	"os"
)

type PlainFile struct {
	path string
}

func (f PlainFile) Write(p []byte) (n int, err error) {
	file, err := os.OpenFile(f.path, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	n, err = file.Write(p)
	return
}

func (f PlainFile) Read() ([]byte, error) {
	return ioutil.ReadFile(f.path)
}

func NewPlainFile(path string) PlainFile {
	return PlainFile{path}
}
