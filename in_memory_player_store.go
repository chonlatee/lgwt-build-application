package poker

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}, m: sync.Mutex{}}
}

type InMemoryPlayerStore struct {
	m     sync.Mutex
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.m.Lock()
	i.store[name]++
	i.m.Unlock()
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
