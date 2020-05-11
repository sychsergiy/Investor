package file

import (
	"encoding/json"
)

type JSON struct {
	File
}

func (f JSON) Create() error {
	err := f.File.Create()
	if err != nil {
		return err
	}
	// write {} to file to keep JSON format valid
	err = f.WriteJSON(make(map[string]interface{}))
	return err
}

func (f JSON) WriteJSON(data interface{}) (err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	_, err = f.File.Write(jsonData)
	return
}

func NewJSON(file File) JSON {
	return JSON{file}
}
