package file

type Reader interface {
	Read() ([]byte, error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
