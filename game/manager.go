package game

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/processor"
	"github.com/awesomegroupidunno/game-server/state"
	"time"
)

type GameManager struct {
	startTime         time.Time
	lastTick          time.Time
	isStarted         bool
	gameState         state.GameState
	commandsToProcess []cmd.GameCommand
	commandFactory    processor.CommandProcessorFactory
}

func (g *GameManager) Start() {
	g.isStarted = true
	g.startTime = time.Now()

}
func (g *GameManager) Pause() {

}
func (g *GameManager) Resume() {

}
func (g *GameManager) AddCommand(c cmd.GameCommand) {

}
func (g *GameManager) TakeState() state.GameState {
	return g.gameState
}
func (g *GameManager) tick() {

}
