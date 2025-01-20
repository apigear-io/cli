
class Counter extends ProxyActor {
    constructor() {
        super("counter", { count: 0 });
    }
    increment() {
        console.log("increment");
        this.count++;
    }
    decrement() {
        console.log("decrement");
        this.count--;
    }
}

const counter = new Counter();



function main() {
    counter.$onProperty("count", function (value) {
        console.log("count changed", value);
    });
    // console.log(JSON.stringify(counter));
    for (let i = 0; i < 5; i++) {
        console.log("incrementing", counter.count);
        counter.increment();
    }
    for (let i = 0; i < 3; i++) {
        counter.decrement();
    }
    console.log("count", counter.count);
    return counter.count;
}
