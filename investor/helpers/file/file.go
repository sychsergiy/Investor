package file

import "os"

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

func (reader Reader) Read(p []byte) (int, error) {
	file, err := os.Open(reader.Path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	n, err := file.Read(p)
	return n, err
}
