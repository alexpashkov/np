package priority_queue

import (
	"github.com/alexpashkov/npuzzle/src/state"
	"github.com/alexpashkov/npuzzle/src/heuristics"
)

type PriorityQueue struct {
	queue []*state.State
	priorityCalc heuristics.Fn
}

func (pq PriorityQueue) Len() int {
	return len(pq.queue)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.priorityCalc(pq.queue[i].Puzzle) > pq.priorityCalc(pq.queue[j].Puzzle)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

func (pq PriorityQueue) Push(x interface{}) {
	item := x.(*state.State)
	pq.queue = append(pq.queue, item)
}

func (pq PriorityQueue) Pop() (x interface{}) {
	x = pq.queue[0]
	pq.queue = pq.queue[1:len(pq.queue)]
	return
}

func New(fn heuristics.Fn) (pq PriorityQueue) {
	pq.queue = make([]*state.State, 0)
	pq.priorityCalc = fn
	return
}
