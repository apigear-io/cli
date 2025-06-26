// Traffic light simulation with cars
const trafficLight = $createService("trafficLight", {
    state: "red",    // red, yellow, green
    carsWaiting: 0
});

const carGenerator = $createService("carGenerator", {
    carsGenerated: 0,
    interval: 2000   // ms between cars
});

const statistics = $createService("statistics", {
    totalCarsPassed: 0,
    averageWaitTime: 0,
    carsWaitingHistory: []
});

// Traffic light methods using natural API
trafficLight.changeState = function () {
    const previousState = this.state;
    switch (this.state) {
        case "red":
            this.state = "green";
            // Let cars pass while green
            while (this.carsWaiting > 0) {
                this.letCarPass();
            }
            break;
        case "green":
            this.state = "yellow";
            break;
        case "yellow":
            this.state = "red";
            break;
    }
    console.log(`Traffic light changed from ${previousState} to ${this.state}`);
    this.emit('stateChanged', previousState, this.state);
}

trafficLight.letCarPass = function () {
    if (this.state === "green" && this.carsWaiting > 0) {
        this.carsWaiting--;
        statistics.recordCarPassed();
        console.log("Car passed through intersection");
        this.emit('carPassed');
    }
}

trafficLight.addWaitingCar = function () {
    this.carsWaiting++;
    statistics.recordWaitingCar(this.carsWaiting);
}

// Car generator methods using natural API
carGenerator.generateCar = function () {
    this.carsGenerated++;
    trafficLight.addWaitingCar();
    console.log(`Generated car #${this.carsGenerated}`);
    this.emit('carGenerated', this.carsGenerated);
}

// Statistics methods using natural API
statistics.recordCarPassed = function () {
    this.totalCarsPassed++;
}

statistics.recordWaitingCar = function (currentWaiting) {
    this.carsWaitingHistory.push({
        timestamp: Date.now(),
        count: currentWaiting
    });

    // Calculate average waiting time
    if (this.carsWaitingHistory.length > 1) {
        const totalWaitTime = this.carsWaitingHistory.reduce((sum, record, index, array) => {
            if (index === 0) return sum;
            return sum + (record.timestamp - array[index - 1].timestamp);
        }, 0);
        this.averageWaitTime = totalWaitTime / this.totalCarsPassed;
    }
}

function main() {
    // Set up monitoring using natural API
    trafficLight.on("state", function (state) {
        console.log(`Traffic light state changed to: ${state}`);
    });

    trafficLight.on("carsWaiting", function (count) {
        console.log(`Cars waiting: ${count}`);
    });

    statistics.on("totalCarsPassed", function (total) {
        console.log(`Total cars passed: ${total}`);
    });
    
    // Listen to custom signals
    trafficLight.on('stateChanged', function(from, to) {
        console.log(`Light transitioned: ${from} → ${to}`);
    });
    
    trafficLight.on('carPassed', function() {
        console.log('Car passed signal received');
    });

    // Run simulation
    for (let i = 0; i < 5; i++) {
        carGenerator.generateCar();
        trafficLight.changeState(); // Cycle through states
    }

    return {
        carsGenerated: carGenerator.carsGenerated,
        carsPassed: statistics.totalCarsPassed,
        averageWaitTime: statistics.averageWaitTime
    };
}
