# diffview

I did a lot of testing recently where I had to compare complex structs and or strings. I really like the GoLand diff viewer. I really hated to copy stuff from the console to two files in order to use the GoLand diff view. So here is diffview.

## usage

```go

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
```

# contributing

Feel free to add different Openers for different operating systems/diff view tools.
