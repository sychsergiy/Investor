package repositories

import (
	"encoding/json"
	"io"
)

type JsonFileRepository struct {
	recordsReader RecordsJsonReader
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
	reader      io.Reader
	unmarshaler RecordUnmarshaler
}

func (w RecordsJsonReader) Read() (map[string]Record, error) {
	var content []byte
	_, err := w.reader.Read(content)
	if err != nil {
		return nil, err
	}

	recordsMap, err := w.unmarshaler.Unmarshal(content)
	return recordsMap, err
}

type RecordUnmarshaler interface {
	Unmarshal([]byte) (map[string]Record, error)
}

func NewJsonFileRepository(reader RecordsJsonReader, writer RecordsJsonWriter) JsonFileRepository {
	return JsonFileRepository{reader, writer, NewInMemoryRepository()}
}
