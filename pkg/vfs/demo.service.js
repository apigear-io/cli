const counter = $createService("counter", {
    count: 0,
});

counter.onProperty("count", function (count) {
    console.log("count:", count);
});

counter.increment = function () {
    this.count++;
};

counter.reset = function () {
    this.count = 0;
    this.$emit("reset", this.count);
}

function main() {
    setInterval(() => {
        counter.increment();
        if (counter.count >= 0) {
            counter.reset();
        }
    }, 1000);
}

