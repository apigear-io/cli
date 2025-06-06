// create a simulation service
const service = $createService("demo.Counter", {
    count: 0,
});

service.onProperty("count", function (count) {
    console.log("count:", count);
});

service.setMethod("increment", function () {
    console.log("incremented");
    const count = service.getProperty("count");
    if (count >= 10) {
        service.callMethod("reset");
    }
    service.setProperty("count", count + 1);
});

service.setMethod("reset", function () {
    service.setProperty("count", 0);
    service.emitSignal("reset", service.getProperty("count"));
})


// create a simulation client to trigger our own service
const client = $createClient("demo.Counter");

client.onProperty("count", function (count) {
    console.log("count:", count);
});

client.onSignal("reset", function (count) {
    console.log("counter was reset to", count);
});



function main() {
    client.setProperty("count", 1);
    setInterval(() => {
        client.callMethod("increment");
    }, 1000);
}