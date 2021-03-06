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
	startTime          time.Time
	lastTick           time.Time
	lastPowerupDespawn time.Time
	restartTime        time.Time
	gameRestart        int
	isStarted          bool
	isPaused           bool
	gameState          state.GameState
	physicsManager     *processor.Physics
	commandsToProcess  []*cmd.GameCommand
	commandFactory     *processor.CommandProcessorFactory
	responses          chan state.StateResponse
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
	g.restartTime = time.Now()
	g.gameRestart = 5
	g.isStarted = true
	g.isPaused = false
	g.startTime = time.Now()
	g.lastTick = time.Now()
	g.lastPowerupDespawn = time.Now()

	commandsMutex.Unlock()
	stateMutex.Unlock()

	lastGameOverTick := time.Now()
	for {
		g.restartIfNeeded()
		if g.shouldTick() {
			g.tick()
		}
		if g.gameState.GameOver != -1 {
			if time.Now().After(lastGameOverTick.Add(time.Second)) {
				g.gameState.SecToRestart--
				lastGameOverTick = time.Now()
			}
		} else {
			lastGameOverTick = time.Now()
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
	return g.isStarted && !g.isPaused && g.gameState.GameOver == -1
}

//
func (g *GameManager) restartIfNeeded() {
	stateMutex.Lock()
	defer stateMutex.Unlock()
	if g.gameState.GameOver != -1 {
		if time.Now().After(g.restartTime) {
			g.gameState.GameOver = -1
			st := g.physicsManager.NewGameState()
			st.Vehicles = g.gameState.Vehicles
			for _, v := range st.Vehicles {
				v.IsAlive = false
				v.TimeDestroyed = time.Now()
			}
			g.gameState = st
		}
	}
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

	g.physicsManager.ApplyGravity(&g.gameState)

	if time.Since(g.lastPowerupDespawn) >= g.physicsManager.PowerupRespawn && len(g.gameState.PowerUps) < g.physicsManager.MaxPowerups {
		g.physicsManager.SpawnPowerup(&g.gameState)
		g.lastPowerupDespawn = time.Now()
	}

	for _, rocket := range g.gameState.Rockets {
		g.physicsManager.MoveRocket(rocket, tickDuration)
	}

	for _, well := range g.gameState.GravityWells {
		if time.Now().After(well.Expires) {
			well.ShouldRemove = true
		}
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
					g.gameState.GameOver = base.TeamId
					g.restartTime = time.Now().Add(time.Duration(g.gameRestart) * time.Second)
					g.gameState.SecToRestart = g.gameRestart
				}
			}
		}

	}

	for z, vehicle := range g.gameState.Vehicles {
		g.physicsManager.RespawnVehicle(vehicle, g.gameState)
		g.physicsManager.VehicleBounding(vehicle)
		g.physicsManager.MoveVehicle(vehicle, tickDuration)
		g.physicsManager.VehicleFrictionSlow(vehicle, tickDuration)

		for i := z + 1; i < len(g.gameState.Vehicles); i++ {
			if collision.Collides(vehicle, g.gameState.Vehicles[i]) {
				g.physicsManager.VehicleCollision(vehicle, g.gameState.Vehicles[i])
			}
		}
		for _, bullet := range g.gameState.Bullets {
			if bullet.OwnerId != vehicle.Owner && vehicle.IsAlive {
				if collision.Collides(vehicle, bullet) {
					g.physicsManager.DamageVehicle(vehicle, bullet)
				}
			}
		}

		for _, shield := range g.gameState.Shields {
			if shield.IsEnabled {
				if collision.Collides(shield, vehicle) {
					g.physicsManager.VehicleCollision(vehicle, &state.Vehicle{Angle: vehicle.Angle, IsAlive: true})
				}
			}
		}

		for _, shieldGenerator := range g.gameState.ShieldGenerators {
			if collision.Collides(shieldGenerator, vehicle) {
				g.physicsManager.VehicleCollision(vehicle, &state.Vehicle{Angle: vehicle.Angle, IsAlive: true})
			}
		}

		for _, base := range g.gameState.Bases {
			if collision.Collides(base, vehicle) {
				g.physicsManager.VehicleCollision(vehicle, &state.Vehicle{Angle: vehicle.Angle, IsAlive: true})
			}
		}

		for _, powerup := range g.gameState.PowerUps {
			if collision.Collides(powerup, vehicle) {
				g.physicsManager.PickupPowerUp(vehicle, powerup)
				g.lastPowerupDespawn = time.Now()
			}
		}

	}

	g.gameState.Bullets = processor.CleanupBullets(g.gameState.Bullets)
	g.gameState.PowerUps = processor.CleanupPowerups(g.gameState.PowerUps)
	g.gameState.Rockets = processor.CleanupRockets(g.gameState.Rockets)
	g.gameState.GravityWells = processor.CleanupWells(g.gameState.GravityWells)

	stateMutex.Unlock()

}
