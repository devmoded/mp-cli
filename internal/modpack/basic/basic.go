package basic

type BasicModpack struct {
	path string
}

func NewBasicModpack(path string) BasicModpack {
	return BasicModpack{
		path: path,
	}
}
