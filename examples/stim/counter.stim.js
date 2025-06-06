const stim = require("stim")
const counter = stim.createClient("counter")
var count = 0
setInterval(function() {
    count++
    console.log("incrementing")
    counter.setProperty("count", count)
    counter.invoke("increment")
}, 1000)
// const counter = node.object("counter")

// counter.onProperty("count", function (value) {
//     console.log("count changed", value);
// })

// counter.set("count", 10)
// counter.call("increment")
// counter.get("count")
