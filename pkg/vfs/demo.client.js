const node = connect()
const counter = node.link("counter");

counter.onProperty("count", function (count) {
    console.log("count:", count);
});

counter.onSignal("reset", function (count) {
    console.log("counter was reset to", count);
});

function main() {
    counter.setProperty("count", 1);
    setInterval(() => {
        counter.invoke("increment");
    }, 1000);
}