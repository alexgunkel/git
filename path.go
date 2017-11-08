package git

type path struct {
	c string
}

func (p path) add(s string) path {
	return path{p.c + s}
}

func (p path) String() string {
	return p.c
}
