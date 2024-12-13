package geecache

type ByteView struct {
	b []byte
}

func (bv ByteView) Len() int {
	return len(bv.b)
}

func CloneBytes(b []byte) []byte {
    c := make([]byte, len(b))
    copy(c, b)
    return c
}

func (bv ByteView)String() string{
    return string(bv.b)
}
