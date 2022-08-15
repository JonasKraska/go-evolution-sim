package main

import (
	"testing"
)

type Player struct {
	health int
}

func (player *Player) changeHealthReference() {
	player.health += 1
}

func (player Player) changeHealthValue() Player {
	player.health += 1

	return player
}

func UseReference(iter int) *Player {
	player := &Player{health: 0}

	for i := 0; i < iter; i++ {
		player.changeHealthReference()
	}

	return player
}

func UseValue(iter int) Player {
	player := Player{health: 0}

	for i := 0; i < iter; i++ {
		player = player.changeHealthValue()
	}

	return player
}

func Benchmark_UseReference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseReference(10000)
	}
}

func Benchmark_UseValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseValue(10000)
	}
}
