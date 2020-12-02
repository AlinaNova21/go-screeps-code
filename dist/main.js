console.log(`=== New Global ${Game.time} ===`)
require('./encoding')
require('./wasm_exec')

const goMod = require('./go')
const mod = new WebAssembly.Module(goMod)
const go = new Go()
const instance = new WebAssembly.Instance(mod, go.importObject)
go.run(instance)

module.exports = {
  loop () {
    instance.exports.loop()
    console.log(`==========`)
    console.log(`CPU: Used: ${Game.cpu.getUsed().toFixed(3)} Limit: ${Game.cpu.limit} Bucket: ${Game.cpu.bucket}`)
    console.log(`MEMORY: Used: ${(RawMemory.get().length / 1024).toFixed(3)}KB`)
    try {
      // eslint-disable-next-line camelcase
      const { used_heap_size, heap_size_limit, total_available_size } = Game.cpu.getHeapStatistics()
      const MB = (v) => ((v / 1024) / 1024).toFixed(3)
      console.log(`HEAP: Used: ${MB(used_heap_size)}MB Available: ${MB(total_available_size)}MB Limit: ${MB(heap_size_limit)}MB`)
    } catch(e) {
      this.console.log('HEAP: Unavailable')
    }
  }
}