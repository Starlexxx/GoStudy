package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	finalScore  = 21
	finishDiff  = 2
	smashChance = 0.2
)

type PingPong struct {
	score map[string]int
	ch    chan string
	wg    *sync.WaitGroup
}

type Player struct {
	name   string
	server bool
}

func main() {
	p := NewPingPong()
	p1 := NewPlayer()
	p2 := NewPlayer()
	p.Run(p1, p2)
}

func NewPingPong() *PingPong {
	return &PingPong{
		score: make(map[string]int),
		ch:    make(chan string),
		wg:    &sync.WaitGroup{},
	}
}

func NewPlayer() Player {
	fmt.Println("Enter player name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Player", name, "is created!")

	return Player{name: name, server: false}
}

func (p *PingPong) Run(p1, p2 Player) {
	p.score[p1.name] = 0
	p.score[p2.name] = 0
	p.wg.Add(2)
	go p.play(&p1)
	go p.play(&p2)
	p.ch <- "begin"
	p.wg.Wait()
}

func (p *PingPong) play(player *Player) {
	defer p.wg.Done()
	for cmd := range p.ch {
		switch cmd {
		case "begin":
			player.server = true
			p.hit(player)
		case "stop":
			if p.isFinished() {
				close(p.ch)
				for name, score := range p.score {
					fmt.Println(name, "score:", score)
				}
				return
			}
			player.server = !player.server
			fmt.Println(player.name, "scores!")
			p.ch <- "begin"
		default:
			p.hit(player)
		}
	}
}

func (p *PingPong) isFinished() bool {
	for _, score := range p.score {
		if score >= finalScore && p.isWin(score) {
			return true
		}
	}
	return false
}

func (p *PingPong) isWin(score int) bool {
	for _, otherScore := range p.score {
		if score-otherScore >= finishDiff {
			return true
		}
	}
	return false
}

func (p *PingPong) hit(player *Player) {
	if player.server {
		fmt.Println(player.name, ": ping")
		if p.isSmash() {
			p.score[player.name]++
			p.ch <- "stop"
		} else {
			p.ch <- "ping"
		}
	} else {
		fmt.Println(player.name, ": pong")
		if p.isSmash() {
			p.score[player.name]++
			p.ch <- "stop"
		} else {
			p.ch <- "pong"
		}
	}
	time.Sleep(100 * time.Millisecond)
}

func (p *PingPong) isSmash() bool {
	return rand.Float64() < smashChance
}
