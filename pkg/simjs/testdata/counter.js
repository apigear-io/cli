const counter = $world.createActor('counter', { count: 0 })
let count = 0
counter.onChange('count', (v) => count = v)
counter.onMethod('increment', () => counter.set('count', counter.get('count') + 1))

function main() {
    return count
}

function add(a, b) {
    return a + b
}