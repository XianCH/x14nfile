package x14nfile

type Tree interface {
	Insert(key string, value any)
	Search(key string) (any, error)
	Delete(key string) error
	Display() error
}

type TreeError struct {
	Code    int
	Message string
}

// b tree

// b plus tree
