call("demo/Counter#increment", { step: 5})
set("demo/Counter", { count: 1})
signal("demo/Counter#shutdown", { time: 1000 })
