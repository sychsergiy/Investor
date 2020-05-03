package file

type JsonWriter interface {
	WriteJson(interface{}) error
}

type Reader interface {
	Read() ([]byte, error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Creator interface {
	Create() error
}

type Exists interface {
	Exists() (bool, error)
}

type CreatorExists interface {
	Creator
	Exists
}

type File interface {
	Reader
	Writer
	CreatorExists
}
