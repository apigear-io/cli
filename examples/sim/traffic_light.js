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

// Traffic light methods
trafficLight.changeState = function () {
    switch (trafficLight.state) {
        case "red":
            trafficLight.state = "green";
            // Let cars pass while green
            while (trafficLight.carsWaiting > 0) {
                trafficLight.letCarPass();
            }
            break;
        case "green":
            trafficLight.state = "yellow";
            break;
        case "yellow":
            trafficLight.state = "red";
            break;
    }
    console.log(`Traffic light changed to ${trafficLight.state}`);
}

trafficLight.letCarPass = function () {
    if (trafficLight.state === "green" && trafficLight.carsWaiting > 0) {
        trafficLight.carsWaiting--;
        statistics.recordCarPassed();
        console.log("Car passed through intersection");
    }
}

trafficLight.addWaitingCar = function () {
    trafficLight.carsWaiting++;
    statistics.recordWaitingCar(trafficLight.carsWaiting);
}

// Car generator methods
carGenerator.generateCar = function () {
    carGenerator.carsGenerated++;
    trafficLight.addWaitingCar();
    console.log(`Generated car #${carGenerator.carsGenerated}`);
}

// Statistics methods
statistics.recordCarPassed = function () {
    statistics.totalCarsPassed++;
}

statistics.recordWaitingCar = function (currentWaiting) {
    statistics.carsWaitingHistory.push({
        timestamp: Date.now(),
        count: currentWaiting
    });

    // Calculate average waiting time
    if (statistics.carsWaitingHistory.length > 1) {
        const totalWaitTime = statistics.carsWaitingHistory.reduce((sum, record, index, array) => {
            if (index === 0) return sum;
            return sum + (record.timestamp - array[index - 1].timestamp);
        }, 0);
        statistics.averageWaitTime = totalWaitTime / statistics.totalCarsPassed;
    }
}

function main() {
    // Set up monitoring
    trafficLight.$.onProperty("state", function (state) {
        console.log(`Traffic light state changed to: ${state}`);
    });

    trafficLight.$.onProperty("carsWaiting", function (count) {
        console.log(`Cars waiting: ${count}`);
    });

    statistics.$.onProperty("totalCarsPassed", function (total) {
        console.log(`Total cars passed: ${total}`);
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
