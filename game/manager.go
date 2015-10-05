package game

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/processor"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
	"sync"
	"time"
)

var stateMutex sync.Mutex
var commandsMutex sync.Mutex

type GameManager struct {
	startTime         time.Time
	lastTick          time.Time
	isStarted         bool
	isPaused          bool
	gameState         state.GameState
	commandsToProcess []*cmd.GameCommand
	commandFactory    processor.CommandProcessorFactory
	Responses         chan state.StateResponse
}

func (g *GameManager) Start() {
	stateMutex.Lock()
	commandsMutex.Lock()

	g.commandsToProcess = []*cmd.GameCommand{}
	g.gameState = state.GameState{}
	g.isStarted = true
	g.isPaused = false
	g.startTime = time.Now()

	commandsMutex.Unlock()
	stateMutex.Unlock()

	for g.isStarted && !g.isPaused && !g.gameState.GameOver {
		g.tick()
		time.Sleep(5 * time.Millisecond)
		g.Responses <- state.StateResponse{State: g.gameState}
	}

}
func (g *GameManager) Pause() {
	stateMutex.Lock()
	g.isPaused = true
	stateMutex.Unlock()
}
func (g *GameManager) Resume() {
	stateMutex.Lock()
	g.isPaused = false
	stateMutex.Unlock()
}
func (g *GameManager) AddCommand(c cmd.GameCommand) {
	commandsMutex.Lock()
	g.commandsToProcess = append(g.commandsToProcess, &c)
	commandsMutex.Unlock()
}
func (g *GameManager) TakeState() state.GameState {
	stateMutex.Lock()
	a := g.gameState.Copy()
	stateMutex.Unlock()
	return a
}
func (g *GameManager) tick() {
	commandsMutex.Lock()
	commands := g.commandsToProcess
	g.commandsToProcess = g.commandsToProcess[:0]
	commandsMutex.Unlock()

	if len(commands) > 0 {
		log.Printf("Ticking with %d commands", len(commands))
	}

	stateMutex.Lock()
	stateMutex.Unlock()

}
