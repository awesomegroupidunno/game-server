package game

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/processor"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
	"sync"
	"time"
)

type GameManager struct {
	startTime         time.Time
	lastTick          time.Time
	isStarted         bool
	isPaused          bool
	stateMutex        *sync.Mutex
	commandMutex      *sync.Mutex
	gameState         state.GameState
	commandsToProcess []*cmd.GameCommand
	commandFactory    processor.CommandProcessorFactory
}

func (g *GameManager) Start() {
	g.stateMutex = &sync.Mutex{}
	g.commandMutex = &sync.Mutex{}

	g.commandMutex.Lock()
	g.commandsToProcess = []*cmd.GameCommand{}
	g.commandMutex.Unlock()

	g.stateMutex.Lock()
	g.gameState = state.GameState{}
	g.stateMutex.Unlock()

	g.isStarted = true
	g.isPaused = false
	g.startTime = time.Now()

	for g.isStarted && !g.isPaused && !g.gameState.GameOver {
		g.tick()
		time.Sleep(5 * time.Millisecond)
	}

}
func (g *GameManager) Pause() {
	g.stateMutex.Lock()
	g.isPaused = true
	g.stateMutex.Unlock()
}
func (g *GameManager) Resume() {
	g.stateMutex.Lock()
	g.isPaused = false
	g.stateMutex.Unlock()
}
func (g *GameManager) AddCommand(c cmd.GameCommand) {
	g.commandMutex.Lock()
	g.commandsToProcess = append(g.commandsToProcess, &c)
	g.commandMutex.Unlock()
}
func (g *GameManager) TakeState() state.GameState {
	g.stateMutex.Lock()
	a := g.gameState.Copy()
	g.stateMutex.Unlock()
	return a
}
func (g *GameManager) tick() {
	g.commandMutex.Lock()
	commands := g.commandsToProcess
	g.commandsToProcess = g.commandsToProcess[:0]
	g.commandMutex.Unlock()

	if len(commands) > 0 {
		log.Printf("Ticking with %d commands", len(commands))
	}

	g.stateMutex.Lock()
	g.stateMutex.Unlock()

}
