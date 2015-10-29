package game

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/processor"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
	"sync"
	"time"
)

// used for synchronizing on state variables
var stateMutex sync.Mutex

// used for synchronizing on commands array
// should lock and unlock any time that:
//	g.commandsToProcess is accessed
var commandsMutex sync.Mutex

type GameManager struct {
	startTime         time.Time
	lastTick          time.Time
	isStarted         bool
	isPaused          bool
	gameState         state.GameState
	physicsManager    *processor.Physics
	commandsToProcess []*cmd.GameCommand
	commandFactory    *processor.CommandProcessorFactory
	responses         chan state.StateResponse
}

// returns a new manager for a game
// GameManager is a producer of StateResponses and places produced values in response_channel
// game_state is the starting GameState of the game
//
// 	state.NewGameState()
// creates an empty state
func NewManager(gameState state.GameState, responseChannel chan state.StateResponse, factory *processor.CommandProcessorFactory) GameManager {
	stateMutex.Lock()
	commandsMutex.Lock()
	defer stateMutex.Unlock()
	defer commandsMutex.Unlock()
	return GameManager{gameState: gameState, responses: responseChannel, commandFactory: factory, physicsManager: factory.Physics}
}

// starts the gamemanager
// does not spawn any new goroutines
func (g *GameManager) Start() {
	stateMutex.Lock()
	commandsMutex.Lock()

	g.commandsToProcess = []*cmd.GameCommand{}
	g.isStarted = true
	g.isPaused = false
	g.startTime = time.Now()
	g.lastTick = time.Now()

	commandsMutex.Unlock()
	stateMutex.Unlock()

	for {
		if g.isStarted && !g.isPaused {
			g.tick()
		}
		stateMutex.Lock()
		g.lastTick = time.Now()
		stateMutex.Unlock()

		g.responses <- state.StateResponse{State: g.gameState}
		time.Sleep(15 * time.Millisecond)

	}

}

// Pauses execution of the game_manager
// while paused, it will no longer produce StateResponses
// Threadsafe
func (g *GameManager) Pause() {
	stateMutex.Lock()
	g.isPaused = true
	stateMutex.Unlock()
}

//Resumes execution
// Treadsafe
func (g *GameManager) Resume() {
	stateMutex.Lock()
	g.isPaused = false
	stateMutex.Unlock()
}

// Adds a command to be processed by the GameManager
// Command will be processed on next tick of the GameManager
// Threadsafe
func (g *GameManager) AddCommand(c cmd.GameCommand) {
	commandsMutex.Lock()
	g.commandsToProcess = append(g.commandsToProcess, &c)
	commandsMutex.Unlock()
}

// Performs next step in the game
// Empties g.commandsToProcess and processes them
// Also progresses GameState based upon elapsed time
// Threadsafe
func (g *GameManager) tick() {
	commandsMutex.Lock()
	commands := g.commandsToProcess
	g.commandsToProcess = g.commandsToProcess[:0]
	commandsMutex.Unlock()

	stateMutex.Lock()
	tickDuration := time.Since(g.lastTick)
	if len(commands) > 0 {
		log.Printf("Ticking with %d commands", len(commands))
	}
	for _, command := range commands {
		proc := g.commandFactory.GetCommandProcessor(command)
		proc.Run(&(g.gameState), *command)
	}

	for _, vehicle := range g.gameState.Vehicles {
		g.physicsManager.MoveVehicle(vehicle, tickDuration)
		g.physicsManager.VehicleFrictionSlow(vehicle, tickDuration)
	}

	stateMutex.Unlock()

}
