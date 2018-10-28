package priority_queue

import (
	"container/heap"
	"github.com/alexpashkov/npuzzle/src/heuristics"
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"github.com/alexpashkov/npuzzle/src/state"
	"testing"
)

var manhattanHeuristics = heuristics.Funcs["manhattan"]

func setupPriorityQueue(t *testing.T) (pq PriorityQueue) {
	pq = New(manhattanHeuristics)
	itemsList := []*state.State{
		{
			Parent: nil,
			Puzzle: puzzle.Puzzle{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			Parent: nil,
			Puzzle: puzzle.Puzzle{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}},
		},
	}
	t.Log("The priority for the first item: ", itemsList[0].F(manhattanHeuristics))
	t.Log("The priority for the second item: ", itemsList[1].F(manhattanHeuristics))
	heap.Push(&pq, itemsList[0])
	heap.Push(&pq, itemsList[1])
	return
}

func TestPriorityQueue_Len(t *testing.T) {
	pq := setupPriorityQueue(t)

	if pq.Len() != 2 {
		t.Error("Invalid length of PriorityQueue. Expected: 2. Got: ", pq.Len())
	}
}

func TestPriorityQueue_Has(t *testing.T) {
	pq := PriorityQueue{
		priorityCalc: manhattanHeuristics,
	}
	item := &state.State{
		Parent: nil,
		Puzzle: puzzle.Puzzle{{5, 3, 1}, {2, 4, 6}, {0, 7, 8}},
	}

	pq.queue = append(pq.queue, item)

	if !pq.Has(item) {
		t.Error("Method 'Has' was not able to find an existing item.")
	}
}

func TestPriorityQueue_Has2(t *testing.T) {
	pq := PriorityQueue{
		priorityCalc: manhattanHeuristics,
	}
	item := &state.State{
		Parent: nil,
		Puzzle: puzzle.Puzzle{{5, 3, 1}, {2, 4, 6}, {0, 7, 8}},
	}

	pq.queue = append(pq.queue, item)

	if pq.Has(&state.State{
		Parent: nil,
		Puzzle: puzzle.Puzzle{{2, 1, 3}, {4, 5, 6}, {7, 8, 0}},
	}) {
		t.Error("Method 'Has' has found the element that is not in the queue.")
	}
}

func TestPriorityQueue_Push(t *testing.T) {
	pq := setupPriorityQueue(t)

	item := &state.State{
		Parent: nil,
		Puzzle: puzzle.Puzzle{{5, 3, 1}, {2, 4, 6}, {0, 7, 8}},
	}
	heap.Push(&pq, item)

	if pq.Len() != 3 {
		t.Error("Invalid length of PriorityQueue. Expected: 3. Got: ", pq.Len())
	}

	if !pq.Has(item) {
		t.Error("Push hasn't updated the internal queue.")
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	pq := New(manhattanHeuristics)
	itemList := []*state.State{
		{
			Parent: nil,
			Puzzle: puzzle.Puzzle{{5, 3, 1}, {2, 4, 6}, {0, 7, 8}},
		},
		{
			Parent: nil,
			Puzzle: puzzle.Puzzle{{1, 2, 3}, {4, 5, 6}, {7, 0, 8}},
		},
		{
			Parent: nil,
			Puzzle: puzzle.Puzzle{{3, 1, 2}, {5, 8, 4}, {6, 7, 0}},
		},
		{
			Parent: nil,
			Puzzle: puzzle.Puzzle{{8, 1, 2}, {5, 7, 3}, {0, 4, 6}},
		},
		{
			Parent: nil,
			Puzzle: puzzle.Puzzle{{2, 8, 5}, {4, 0, 7}, {1, 3, 6}},
		},
	}

	for i, item := range itemList {
		t.Log("The priority of the ", i, " item: ", item.F(manhattanHeuristics))
		heap.Push(&pq, item)
	}

	lastPriority := 0
	for i := range itemList {
		item := heap.Pop(&pq).(*state.State)
		currPriority := item.F(manhattanHeuristics)
		t.Log("The priority of the ", i, " popped item: ", currPriority)

		if currPriority < lastPriority {
			t.Error("Item being popped has lower priority value than the one being popped before: ", currPriority, " vs previous ", lastPriority)
		}
		lastPriority = currPriority
	}
}
