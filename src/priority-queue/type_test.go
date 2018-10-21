package priority_queue

import (
	"github.com/alexpashkov/npuzzle/src/heuristics"
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"github.com/alexpashkov/npuzzle/src/state"
	"testing"
)

func setupPriorityQueue() (pq PriorityQueue) {
	pq = New(heuristics.Manhattan)
	pq.Push(&state.State{
		Parent: nil,
		Puzzle: puzzle.Puzzle{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}},
	})
	pq.Push(&state.State{
		Parent: nil,
		Puzzle: puzzle.Puzzle{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}},
	})
	return
}

func TestPriorityQueue_Len(t *testing.T) {
	pq := setupPriorityQueue()

	if pq.Len() != 2 {
		t.Error("Invalid length of PriorityQueue. Expected: 2. Got: ", pq.Len())
	}
}

func TestPriorityQueue_Has(t *testing.T) {
	pq := PriorityQueue{
		priorityCalc: heuristics.Manhattan,
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
		priorityCalc: heuristics.Manhattan,
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
	pq := setupPriorityQueue()

	item := &state.State{
		Parent: nil,
		Puzzle: puzzle.Puzzle{{5, 3, 1}, {2, 4, 6}, {0, 7, 8}},
	}
	pq.Push(item)

	if pq.Len() != 3 {
		t.Error("Invalid length of PriorityQueue. Expected: 3. Got: ", pq.Len())
	}

	if !pq.Has(item) {
		t.Error("Push hasn't updated the internal queue.")
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	pq := PriorityQueue{
		priorityCalc: heuristics.Manhattan,
	}
	itemList := []*state.State{
		{
			Parent: nil,
			Puzzle: puzzle.Puzzle{{5, 3, 1}, {2, 4, 6}, {0, 7, 8}},
		},
		{
			Parent: nil,
			Puzzle: puzzle.Puzzle{{1, 2, 3}, {4, 5, 6}, {7, 0, 8}},
		},
	}

	pq.Push(itemList[1])
	pq.Push(itemList[0])

	poppedItem := pq.Pop()

	if poppedItem != itemList[1] {
		t.Error("Item being popped did not match the expected.")
	}
}
