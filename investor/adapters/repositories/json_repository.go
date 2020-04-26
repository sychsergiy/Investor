package repositories

import (
	"encoding/json"
	"io"
	"os"
)

type FileReader struct {
	Path string
}

type FileWriter struct {
	Path string
}

func (fw FileWriter) Write(p []byte) (n int, err error) {
	f, err := os.Create(fw.Path)
	if err != nil {
		return 0, err
	}
	n, err = f.Write(p)
	return
}

func (reader FileReader) Read(p []byte) (int, error) { // todo: change on io.Reader
	file, err := os.Open(reader.Path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	n, err := file.Read(p)
	return n, err
}

type JsonFileRepository struct {
	recordsReader RecordsJsonReader // todo: change on Json Reader
	recordsWriter RecordsJsonWriter
	repository    InMemoryRepository
}

func (r JsonFileRepository) Create(record Record) (err error) {
	if err = r.repository.Create(record); err != nil {
		return
	} else {
		err = r.dump()
		return
	}
}

func (r JsonFileRepository) CreateBulk(records []Record) (int, error) {
	count, err := r.repository.CreateBulk(records)
	if err != nil {
		return count, err
	}
	err = r.dump()
	return count, err
}

func (r JsonFileRepository) dump() error {
	err := r.recordsWriter.Write(r.repository.records)
	return err
}

type RecordsJsonWriter struct {
	writer io.Writer
}

func (w RecordsJsonWriter) Write(recordsMap map[string]Record) error {
	jsonData, err := json.Marshal(recordsMap)
	if err != nil {
		return err
	}
	_, err = w.writer.Write(jsonData)
	return err
}

type RecordsJsonReader struct {
	reader io.Reader
}

func (w RecordsJsonReader) Read() (map[string]Record, error) {
	var content []byte
	_, err := w.reader.Read(content)
	if err != nil {
		return nil, err
	}

	var recordsMap map[string]Record
	err = json.Unmarshal(content, &recordsMap)
	return recordsMap, err
}

func NewJsonFileRepository() JsonFileRepository {
	return JsonFileRepository{
		RecordsJsonReader{FileReader{"payments.json"}},
		RecordsJsonWriter{FileWriter{"payments.json"}},
		NewInMemoryRepository(),
	}
}
