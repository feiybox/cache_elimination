package cache_elimination

import "container/heap"

type KeyValue struct {
	Key    string
	Value  interface{}
	weight int
	index  int
}

type Queue []*KeyValue

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].weight < q[j].weight
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index, q[j].index = i, j
}

func (q *Queue) Push(v interface{}) {
	en := v.(*KeyValue)
	en.index = q.Len()
	*q = append(*q, en)
}

func (q *Queue) Pop() interface{} {
	oldQue := *q
	n := len(oldQue)
	en := oldQue[n-1]
	oldQue[n-1] = nil
	*q = oldQue[:n-1]
	return en
}

func (q *Queue) update(en *KeyValue, val interface{}, weight int) {
	en.Value = val
	en.weight = weight
	(*q)[en.index] = en
	heap.Fix(q, en.index) // 进行对排序处理
}
