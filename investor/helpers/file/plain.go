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

func (f PlainFile) Create() error {
	_, err := os.Create(f.path)
	return err
}

func (f PlainFile) Exists() (bool, error) {
	info, err := os.Stat(f.path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return !info.IsDir(), nil
}

func NewPlainFile(path string) PlainFile {
	return PlainFile{path}
}

func CreateIfNotExists(file CreatorExists) (bool, error) {
	exists, err := file.Exists()
	if err != nil {
		return false, err
	}
	if exists {
		return false, nil
	} else {
		err := file.Create()
		return true, err
	}
}
