package screeps

type voidFn func()

var setupFns = []voidFn{}
var preLoopFns = []voidFn{}
var postLoopFns = []voidFn{}

func RegisterSetup(fn voidFn) {
	setupFns = append(setupFns, fn)
}

func RegisterPreLoop(fn voidFn) {
	preLoopFns = append(preLoopFns, fn)
}

func RegisterPostLoop(fn voidFn) {
	postLoopFns = append(postLoopFns, fn)
}

func RunSetup() {
	for _, fn := range setupFns {
		fn()
	}
}

func RunPreLoop() {
	for _, fn := range preLoopFns {
		fn()
	}
}

func RunPostLoop() {
	for _, fn := range preLoopFns {
		fn()
	}
}
