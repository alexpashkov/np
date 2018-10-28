package priority_queue

import (
	"container/heap"
	"github.com/alexpashkov/npuzzle/src/heuristics"
	"github.com/alexpashkov/npuzzle/src/state"
	"reflect"
)

type PriorityQueue struct {
	queue        []*state.State
	priorityCalc heuristics.Func
}

func (pq PriorityQueue) Len() int {
	return len(pq.queue)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.queue[i].F(pq.priorityCalc) < pq.queue[j].F(pq.priorityCalc)
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*state.State)
	pq.queue = append(pq.queue, item)
}

func (pq *PriorityQueue) Pop() (x interface{}) {
	x = pq.queue[len(pq.queue)-1]
	pq.queue = pq.queue[0 : len(pq.queue)-1]
	return
}

func (pq PriorityQueue) Has(x interface{}) bool {
	desiredItem := x.(*state.State)

	for _, item := range pq.queue {
		if reflect.DeepEqual(*desiredItem, *item) {
			return true
		}
	}

	return false
}

func New(fn heuristics.Func) (pq PriorityQueue) {
	pq.queue = make([]*state.State, 0)
	pq.priorityCalc = fn
	heap.Init(&pq)
	return
}
