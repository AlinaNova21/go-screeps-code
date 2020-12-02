package screeps

import "syscall/js"

// Game is global.Game
var Game *screepsGame

func init() {
	RegisterSetup(func() {
		Game = initScreepsGame()
	})
}

type screepsGame struct {
	js   js.Value
	CPU  *screepsCPU
	Time int
}

func initScreepsGame() *screepsGame {
	return &screepsGame{
		CPU: initScreepsCPU(),
	}
}

func (g *screepsGame) preLoop() {
	g.js = js.Global().Get("Game")
	g.Time = g.js.Get("time").Int()
	g.CPU.js = g.js.Get("cpu")
	g.CPU.preLoop()
}

func (g *screepsGame) postLoop() {
	g.js = js.Value{}
	g.CPU.js = js.Value{}
}
