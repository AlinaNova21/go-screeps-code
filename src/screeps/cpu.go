package screeps

import "syscall/js"

type heapStats struct {
	UsedHeapSize       int
	HeapSizeLimit      int
	TotalAvailableSize int
}

type screepsCPU struct {
	js        js.Value
	Limit     int
	Bucket    int
	heapStats *heapStats
}

func initScreepsCPU() *screepsCPU {
	return &screepsCPU{
		heapStats: &heapStats{},
	}
}

func (c *screepsCPU) preLoop() {
	c.Limit = c.js.Get("limit").Int()
	c.Bucket = c.js.Get("bucket").Int()
}

func (c *screepsCPU) GetUsed() float64 {
	return c.js.Call("getUsed").Float()
}

func (c *screepsCPU) GetHeapStatistics() *heapStats {
	js := c.js.Call("getHeapStatistics")
	c.heapStats.UsedHeapSize = js.Get("used_heap_size").Int()
	c.heapStats.HeapSizeLimit = js.Get("heap_size_limit").Int()
	c.heapStats.TotalAvailableSize = js.Get("total_available_size").Int()
	return c.heapStats
}
