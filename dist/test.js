require('./encoding')
require('./wasm_exec')

let started = false


const goMod = require('fs').readFileSync(__dirname + '/go.wasm')
const mod = new WebAssembly.Module(goMod)
const go = new Go()
const instance = new WebAssembly.Instance(mod, go.importObject)
go.run(instance)

module.exports = {
  loop() {
    console.log('loop')
    instance.exports.loop()
  }
}