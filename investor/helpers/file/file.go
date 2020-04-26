package file

import (
	"io/ioutil"
	"os"
)

type Reader struct {
	Path string
}

type Writer struct {
	Path string
}

func (fw Writer) Write(p []byte) (n int, err error) {
	f, err := os.Create(fw.Path)
	if err != nil {
		return 0, err
	}
	n, err = f.Write(p)
	return
}

func (reader Reader) Read() ([]byte, error) {
	return ioutil.ReadFile(reader.Path)
}
