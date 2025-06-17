// create a simulation service
const service = $createService("demo.Counter", {
    count: 0,
});

service.$.onProperty("count", function (count) {
    console.log("count:", count);
});

service.increment = function() {
    console.log("incremented");
    const count = service.count
    if (count >= 10) {
        service.reset()
    }
    service.count++
}

service.reset = function() {
    service.count = 0
    service.$.emitSignal("reset", service.getProperty("count"));
}


function main() {
    service.increment()
}