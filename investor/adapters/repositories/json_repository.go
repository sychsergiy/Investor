package repositories

import (
	"encoding/json"
	"investor/helpers/file"
	"io"
	"log"
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

func (r JsonFileRepository) restore() error {
	// todo: create empty dict json file on init
	recordsMap, err := r.recordsReader.Read() // todo: change on records list

	var records []Record
	for _, value := range recordsMap {
		records = append(records, value)
	}

	if err != nil {
		return err
	} else {
		_, err := r.repository.CreateBulk(records)
		if err != nil {
			return err
		}
	}
	return nil
}

type RecordsJsonWriter struct {
	writer io.Writer
}

func NewRecordsJsonWriter(writer io.Writer) RecordsJsonWriter {
	return RecordsJsonWriter{writer}
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
	reader      file.Reader
	unmarshaler RecordUnmarshaler
}

func NewRecordsJsonReader(reader file.Reader, unmarshaler RecordUnmarshaler) RecordsJsonReader {
	return RecordsJsonReader{reader, unmarshaler}
}

func (w RecordsJsonReader) Read() (map[string]Record, error) {
	content, err := w.reader.Read()
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
	repo := JsonFileRepository{reader, writer, NewInMemoryRepository()}
	err := repo.restore()
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
