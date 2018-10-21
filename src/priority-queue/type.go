package priority_queue

import (
	"github.com/alexpashkov/npuzzle/src/state"
	"github.com/alexpashkov/npuzzle/src/heuristics"
	"reflect"
)

type PriorityQueue struct {
	queue        []*state.State
	priorityCalc heuristics.Fn
}

func (pq PriorityQueue) Len() int {
	return len(pq.queue)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.queue[i].H(pq.priorityCalc) > pq.queue[j].H(pq.priorityCalc)
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*state.State)
	pq.queue = append(pq.queue, item)
}

func (pq *PriorityQueue) Pop() (x interface{}) {
	x = pq.queue[0]
	pq.queue = pq.queue[1:len(pq.queue)]
	return
}

func (pq PriorityQueue) Has(x interface{}) bool {
	desiredItem := x.(*state.State)

	for _, item := range pq.queue {
		if reflect.DeepEqual(*desiredItem, *item) {
			return true;
		}
	}

	return false
}

func New(fn heuristics.Fn) (pq PriorityQueue) {
	pq.queue = make([]*state.State, 0)
	pq.priorityCalc = fn
	return
}
