const counter = $createActor("counter", { count: 0 });

counter.$setState({
    count: 10,
});

counter.increment = function () {
    this.count++;
}

counter.decrement = function () {
    this.count--;
}

function main() {
    counter.$onProperty("count", function (value) {
        console.log("count changed", value);
    });
    // console.log(JSON.stringify(counter));
    for (let i = 0; i < 5; i++) {
        counter.increment();
    }
    for (let i = 0; i < 3; i++) {
        counter.decrement();
    }
    console.log("count", counter.count);
    return counter.count;
}
