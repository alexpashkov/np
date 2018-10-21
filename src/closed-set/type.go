package closed_set

import "github.com/alexpashkov/npuzzle/src/state"

type Closed map[state.Id]state.State

func (set Closed) Add(s state.State)  {
	set[s.Id()] = s
}

func (set Closed) Has(s state.State) bool {
	_, has := set[s.Id()]
	return has
}
