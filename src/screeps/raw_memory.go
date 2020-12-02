package screeps

import "syscall/js"

// RawMemory is global.RawMemory
var RawMemory *screepsRawMemory

func init() {
	RegisterSetup(func() {
		RawMemory = &screepsRawMemory{}
		RawMemory.CPU = &screepsCPU{}
	})
}

type screepsRawMemory struct {
	js   js.Value
	CPU  *screepsCPU
	Time int
}

func (r *screepsRawMemory) preLoop() {
	r.js = js.Global().Get("RawMemory")
}

func (r *screepsRawMemory) Get() string {
	return r.js.Call("get").String()
}

func (r *screepsRawMemory) Set(v string) {
	r.js.Call("get", v)
}
