package diffview

import "testing"

var (
	a = `foo



bar



baz
`
	b = `foo



baz



bar`
)

func TestGolandDiffView(t *testing.T) {
	NewGoland().DiffViewBytes("test", []byte(a), []byte(b))
}
