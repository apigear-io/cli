// Connect to ObjectLink backend
const client = connect('ws://localhost:5550/ws');
const counter = client.interface("demo.Counter")
counter.invoke("increment")

