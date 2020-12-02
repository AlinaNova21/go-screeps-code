package main

import (
	"fmt"

	"github.com/ags131/go-screeps-code/src/screeps"
)

func main() {
	fmt.Println("Hello from Go!")
	screeps.RunSetup()
}

//export loop
func loop() {
	screeps.PreLoop()
	fmt.Printf("Loop from Go! Time: %d\n", screeps.Game.Time)
	// cpu := screeps.Game.CPU
	// fmt.Printf("CPU: Used: %.3f Limit: %d Bucket %d\n", cpu.GetUsed(), cpu.Limit, cpu.Bucket)

	// memUsed := len(screeps.RawMemory.Get())
	// fmt.Printf("MEMORY: Used: %.3fKB\n", float64(memUsed)/1024)

	// heapStats := cpu.GetHeapStatistics()
	// fmt.Printf("HEAP: Used: %.3fMB Available: %.3fMB Limit %.3fMB\n", mb(float64(heapStats.UsedHeapSize)), mb(float64(heapStats.TotalAvailableSize)), mb(float64(heapStats.HeapSizeLimit)))
	screeps.PostLoop()
}

func mb(v float64) float64 {
	return (v / 1024) / 1024
}
