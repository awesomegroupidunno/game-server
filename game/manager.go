package game

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/processor"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
	"sync"
	"time"
)

var managerLock sync.Mutex

type GameManager struct {
	startTime         time.Time
	lastTick          time.Time
	isStarted         bool
	isPaused          bool
	gameState         state.GameState
	commandsToProcess []*cmd.GameCommand
	commandFactory    processor.CommandProcessorFactory
}

func (g *GameManager) Start() {
	managerLock.Lock()

	g.commandsToProcess = []*cmd.GameCommand{}
	g.gameState = state.GameState{}
	g.isStarted = true
	g.isPaused = false
	g.startTime = time.Now()

	managerLock.Unlock()

	for g.isStarted && !g.isPaused && !g.gameState.GameOver {
		g.tick()
		time.Sleep(5 * time.Millisecond)
	}

}
func (g *GameManager) Pause() {
	managerLock.Lock()
	g.isPaused = true
	managerLock.Unlock()
}
func (g *GameManager) Resume() {
	managerLock.Lock()
	g.isPaused = false
	managerLock.Unlock()
}
func (g *GameManager) AddCommand(c cmd.GameCommand) {
	managerLock.Lock()
	g.commandsToProcess = append(g.commandsToProcess, &c)
	managerLock.Unlock()
}
func (g *GameManager) TakeState() state.GameState {
	managerLock.Lock()
	a := g.gameState.Copy()
	managerLock.Unlock()
	return a
}
func (g *GameManager) tick() {
	managerLock.Lock()
	commands := g.commandsToProcess
	g.commandsToProcess = g.commandsToProcess[:0]
	managerLock.Unlock()

	if len(commands) > 0 {
		log.Printf("Ticking with %d commands", len(commands))
	}

}
