package heap

// Heap 堆结构，包括一个slice，slice的初始化容量，堆节点个数
type Heap struct {
	a     []int
	n     int
	count int
}

//NewHeap init heap
func NewHeap(capacity int) *Heap {
	return &Heap{
		a: make([]int, capacity+1),
		n: capacity,
	}
}

// Insert 插入 大顶堆 从下往上的堆化方法。 top-max heap -> heapify from down to up
func (h *Heap) Insert(v int) {
	//defensive
	if h.count == h.n {
		return
	}

	h.count++
	i := h.count
	// 把值放到堆最后
	h.a[i] = v

	//compare with parent node
	parent := i / 2
	for parent > 0 && h.a[parent] < h.a[i] {
		h.swap(parent, i)
		i = parent
		parent = i / 2
	}
}

//RemoveMax 大顶堆 heapify from up to down
func (h *Heap) RemoveMax() {
	//defensive
	if h.count == 0 {
		return
	}

	//swap max and last
	h.swap(1, h.count)
	// 交换之后归0
	h.a[h.count] = 0

	h.count--

	//heapify from up to down
	h.HeapifyUpToDown()
}

//HeapifyUpToDown heapify
func (h *Heap) HeapifyUpToDown() {
	for i := 1; i <= h.count/2; {
		maxIndex := h.threeMax(i)
		if maxIndex == i {
			break
		}

		h.swap(i, maxIndex)
		i = maxIndex
	}

}

// 堆化过程中比较 节点及左右子节点的大小
func (h *Heap) threeMax(i int) (maxIndex int) {
	maxIndex = i
	if h.a[i] < h.a[i*2] {
		maxIndex = i * 2
	}

	if i*2+1 <= h.count && h.a[maxIndex] < h.a[i*2+1] {
		maxIndex = i*2 + 1
	}
	return maxIndex
}

//swap two elements
func (h *Heap) swap(i int, j int) {
	h.a[i], h.a[j] = h.a[j], h.a[i]
}

//Sort sort by ascend, a index begin from 1, has n elements
func (h *Heap) Sort() {
	//defensive
	if h.count == 0 {
		return
	}

	k := h.count
	for h.count >= 1 {
		h.swap(1, h.count)
		h.count--
		h.HeapifyUpToDown()
	}
	h.count = k
}
