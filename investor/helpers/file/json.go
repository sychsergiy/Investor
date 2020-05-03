package file

import (
	"encoding/json"
)

type JsonFile struct {
	File
}

func (f JsonFile) Create() error {
	err := f.File.Create()
	if err != nil {
		return err
	}
	// write {} to file to keep JSON format valid
	err = f.WriteJson(make(map[string]interface{}))
	return err
}

func (f JsonFile) WriteJson(data interface{}) (err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	_, err = f.File.Write(jsonData)
	return
}

func NewJsonFile(file File) JsonFile {
	return JsonFile{file}
}
