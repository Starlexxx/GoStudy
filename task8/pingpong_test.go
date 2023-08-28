package main

import (
	"testing"
)

func TestPingPong(t *testing.T) {
	p := NewPingPong()
	p1 := Player{name: "Player 1", server: false}
	p2 := Player{name: "Player 2", server: false}
	p.Run(p1, p2)

	if !p.isFinished() {
		t.Errorf("Game did not finish")
	}

	if p.score[p1.name] == p.score[p2.name] {
		t.Errorf("Game ended in a tie")
	}

	if p.score[p1.name] < finalScore && p.score[p2.name] < finalScore {
		t.Errorf("Game did not reach final score")
	}
}

func TestHit(t *testing.T) {
	p := NewPingPong()
	p1 := Player{name: "Player 1", server: false}
	p2 := Player{name: "Player 2", server: false}
	p.score[p1.name] = 0
	p.score[p2.name] = 0

	p.hit(&p1)
	if p.score[p1.name] != 0 {
		t.Errorf("Player 1 scored without serving")
	}

	p.hit(&p2)
	if p.score[p2.name] != 0 {
		t.Errorf("Player 2 scored without serving")
	}

	p.hit(&p1)
	if p.score[p1.name] != 0 {
		t.Errorf("Player 1 scored without serving")
	}

	p.hit(&p2)
	if p.score[p2.name] != 0 {
		t.Errorf("Player 2 scored without serving")
	}

	p.hit(&p1)
	if p.score[p1.name] != 1 {
		t.Errorf("Player 1 did not score after serving")
	}

	p.hit(&p2)
	if p.score[p2.name] != 1 {
		t.Errorf("Player 2 did not score after serving")
	}
}

func TestIsSmash(t *testing.T) {
	p := NewPingPong()

	smashCount := 0
	for i := 0; i < 1000; i++ {
		if p.isSmash() {
			smashCount++
		}
	}

	if smashCount < 100 || smashCount > 300 {
		t.Errorf("Smash chance is not within expected range")
	}
}
