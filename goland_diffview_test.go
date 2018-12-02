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
	err := NewGoland().DiffViewBytes("test",[]byte(a),[]byte(b))
	if err != nil {
		t.Fatal(err)
	}
}
