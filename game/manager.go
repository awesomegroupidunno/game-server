package game

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/collision"
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
		if g.shouldTick() {
			g.tick()
		}
		stateMutex.Lock()
		g.lastTick = time.Now()
		stateMutex.Unlock()
		g.responses <- state.StateResponse{State: g.gameState.Copy()}

		time.Sleep(15 * time.Millisecond)

	}

}

// Pauses execution of the game_manager
// while paused, it will no longer produce StateResponses
// Threadsafe
func (g *GameManager) Pause() {
	stateMutex.Lock()
	defer stateMutex.Unlock()
	g.isPaused = true
}

// Pauses execution of the game_manager
// while paused, it will no longer produce StateResponses
// Threadsafe
func (g *GameManager) IsPaused() bool {
	stateMutex.Lock()
	defer stateMutex.Unlock()
	return g.isPaused
}

//Resumes execution
// Treadsafe
func (g *GameManager) Resume() {
	stateMutex.Lock()
	defer stateMutex.Unlock()
	g.isPaused = false
}

//Returns if tick should be called
// Treadsafe
func (g *GameManager) shouldTick() bool {
	stateMutex.Lock()
	defer stateMutex.Unlock()
	return g.isStarted && !g.isPaused && !g.gameState.GameOver
}

// Adds a command to be processed by the GameManager
// Command will be processed on next tick of the GameManager
// Threadsafe
func (g *GameManager) AddCommand(c cmd.GameCommand) {
	commandsMutex.Lock()
	defer commandsMutex.Unlock()
	g.commandsToProcess = append(g.commandsToProcess, &c)
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

	if tickDuration > (50 * time.Millisecond) {
		log.Println("Lag spike detected")
	}

	for _, command := range commands {
		proc := g.commandFactory.GetCommandProcessor(command)
		proc.Run(&(g.gameState), *command)
	}

	for _, bullet := range g.gameState.Bullets {
		g.physicsManager.MoveBullet(bullet, tickDuration)
		g.physicsManager.BoundBullet(bullet)
	}

	for _, bullet := range g.gameState.Bullets {

		for _, shield := range g.gameState.Shields {
			if shield.IsEnabled {
				if collision.Collides(bullet, shield) {
					bullet.ShouldRemove = true
				}
			}
		}

		for _, shieldGenerator := range g.gameState.ShieldGenerators {
			if collision.Collides(shieldGenerator, bullet) {
				g.physicsManager.DamageShieldGenerator(bullet, shieldGenerator)
			}
		}

		for _, base := range g.gameState.Bases {
			if collision.Collides(bullet, base) {
				if g.physicsManager.DamageBase(bullet, base) {
					g.gameState.GameOver = true
				}
			}
		}

	}

	for z, vehicle := range g.gameState.Vehicles {
		g.physicsManager.RespawnVehicle(vehicle)
		g.physicsManager.VehicleBounding(vehicle)
		g.physicsManager.MoveVehicle(vehicle, tickDuration)
		g.physicsManager.VehicleFrictionSlow(vehicle, tickDuration)

		for i := z + 1; i < len(g.gameState.Vehicles); i++ {
			if collision.Collides(vehicle, g.gameState.Vehicles[i]) {
				g.physicsManager.VehicleCollisionPhysics(vehicle, g.gameState.Vehicles[i])
			}
		}
		for _, bullet := range g.gameState.Bullets {
			if bullet.OwnerId != vehicle.Owner {
				if collision.Collides(vehicle, bullet) {
					g.physicsManager.DamageVehicle(vehicle, bullet)
				}
			}
		}

		for _, shield := range g.gameState.Shields {
			if shield.IsEnabled {
				if collision.Collides(shield, vehicle) {
					g.physicsManager.VehicleCollisionPhysics(vehicle, &state.Vehicle{Angle: vehicle.Angle, IsAlive: true})
				}
			}
		}

		for _, shieldGenerator := range g.gameState.ShieldGenerators {
			if collision.Collides(shieldGenerator, vehicle) {
				g.physicsManager.VehicleCollisionPhysics(vehicle, &state.Vehicle{Angle: vehicle.Angle, IsAlive: true})
			}
		}

		for _, base := range g.gameState.Bases {
			if collision.Collides(base, vehicle) {
				g.physicsManager.VehicleCollisionPhysics(vehicle, &state.Vehicle{Angle: vehicle.Angle, IsAlive: true})
			}
		}

	}

	g.gameState.Bullets = processor.CleanupBullets(g.gameState.Bullets)

	stateMutex.Unlock()

}
