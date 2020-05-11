package file

import (
	"encoding/json"
)

type JSONFile struct {
	File
}

func (f JSONFile) Create() error {
	err := f.File.Create()
	if err != nil {
		return err
	}
	// write {} to file to keep JSON format valid
	err = f.WriteJSON(make(map[string]interface{}))
	return err
}

func (f JSONFile) WriteJSON(data interface{}) (err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	_, err = f.File.Write(jsonData)
	return
}

func NewJSONFile(file File) JSONFile {
	return JSONFile{file}
}
