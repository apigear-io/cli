// Connect to ObjectLink backend
const client = connect('ws://localhost:5550/ws');
const counter = client.interface("demo.Counter")
counter.onPropertyChange('count', (value) => {
  console.log("count:", value)
})
const count = counter.invoke("increment")
console.log(JSON.stringify(count))
console.log("count", JSON.stringify(count))

