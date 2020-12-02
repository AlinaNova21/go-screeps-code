package screeps

// PreLoop inits the state at the start of the tick
func PreLoop() {
	Game.preLoop()
	RawMemory.preLoop()
}

// PostLoop frees the state at the end of the tick
func PostLoop() {
	Game.postLoop()
	// RawMemory.postLoop()
}
