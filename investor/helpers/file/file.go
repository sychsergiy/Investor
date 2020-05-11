package file

type JSONWriter interface {
	WriteJSON(interface{}) error
}

type Path interface {
	Path() string
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
	Path
	Reader
	Writer
	CreatorExists
}

type IJSONFile interface {
	File
	JSONWriter
}
