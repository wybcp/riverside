package bubble

import "testing"

var values = []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}

func TestBubbleIncSort(t *testing.T) {
	t.Log(values)
	BubbleIncSort(values)
	t.Log(values)
}

func BenchmarkBubbleIncSort(b *testing.B)  {
	for index := 0; index < b.N; index++ {
		BubbleIncSort(values)
	}
}
func BenchmarkBubbleDecSort(b *testing.B)  {
	for index := 0; index < b.N; index++ {
		BubbleDecSort(values)
	}
}
func TestBubbleDecSort(t *testing.T) {
	t.Log(values)
	BubbleDecSort(values)
	t.Log(values)
}
