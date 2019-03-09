package byte

import (
	"testing"
)

func TestAddExample(t *testing.T) {
	t.Log(AndExample(1))
	t.Log(AndExample(2))

}
func BenchmarkAddExample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AndExample(i)
	}
}
func BenchmarkDivExample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivExamle(i)
	}
}

//     1000	   1401651 ns/op	     801 B/op	     100 allocs/op

// ok  	riverside/byte	2.588s
