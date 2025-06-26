// Heater control system simulation
const heater = $createService("heater", {
    isOn: false,
    power: 2000, // watts
    temperature: 20.0, // celsius
    maxTemp: 30.0,
    minTemp: 15.0
});

const thermostat = $createService("thermostat", {
    targetTemperature: 22.0,
    tolerance: 0.5,
    mode: 'auto' // 'auto' or 'manual'
});

const tempSensor = $createService("tempSensor", {
    currentTemperature: 20.0,
    updateInterval: 1000, // ms
    lastUpdate: Date.now()
});

// Heater methods using natural API
heater.turnOn = function () {
    if (!this.isOn) {
        this.isOn = true;
        console.log("Heater turned ON");
        this.emit('stateChanged', true);
    }
}

heater.turnOff = function () {
    if (this.isOn) {
        this.isOn = false;
        console.log("Heater turned OFF");
        this.emit('stateChanged', false);
    }
}

heater.updateTemperature = function (deltaTime) {
    if (this.isOn) {
        // Simple temperature increase model
        // Temperature rises faster when difference to max temp is larger
        const heatIncrease = (this.maxTemp - this.temperature) * 0.1;
        this.temperature += heatIncrease * (deltaTime / 1000);
    } else {
        // Natural cooling model
        // Temperature falls faster when difference to ambient temp is larger
        const cooling = (this.temperature - tempSensor.currentTemperature) * 0.05;
        this.temperature -= cooling * (deltaTime / 1000);
    }
}

// Thermostat methods using natural API
thermostat.setTargetTemperature = function (temp) {
    if (temp >= heater.minTemp && temp <= heater.maxTemp) {
        this.targetTemperature = temp;
        console.log(`Target temperature set to ${temp}°C`);
        this.checkTemperature();
    } else {
        console.log(`Temperature ${temp}°C is outside allowed range`);
    }
}

thermostat.checkTemperature = function () {
    const currentTemp = tempSensor.currentTemperature;
    const lowerBound = this.targetTemperature - this.tolerance;
    const upperBound = this.targetTemperature + this.tolerance;

    if (currentTemp < lowerBound) {
        heater.turnOn();
    } else if (currentTemp > upperBound) {
        heater.turnOff();
    }
}

thermostat.setMode = function (newMode) {
    if (newMode === 'auto' || newMode === 'manual') {
        this.mode = newMode;
        console.log(`Thermostat mode set to ${newMode}`);
        if (newMode === 'auto') {
            this.checkTemperature();
        }
    }
}

// Temperature sensor methods using natural API
tempSensor.update = function () {
    const now = Date.now();
    const deltaTime = now - this.lastUpdate;
    this.lastUpdate = now;

    // Update current temperature based on heater's influence
    const heatTransfer = (heater.temperature - this.currentTemperature) * 0.1;
    this.currentTemperature += heatTransfer * (deltaTime / 1000);

    // Add some random fluctuation
    this.currentTemperature += (Math.random() - 0.5) * 0.1;

    console.log(`Current temperature: ${this.currentTemperature.toFixed(1)}°C`);

    if (thermostat.mode === 'auto') {
        thermostat.checkTemperature();
    }
}

function main() {
    // Set up monitoring using natural API
    heater.on("isOn", function (isOn) {
        console.log(`Heater state changed to: ${isOn ? "ON" : "OFF"}`);
    });

    tempSensor.on("currentTemperature", function (temp) {
        console.log(`Temperature sensor reading: ${temp.toFixed(1)}°C`);
    });

    thermostat.on("targetTemperature", function (temp) {
        console.log(`Target temperature changed to: ${temp.toFixed(1)}°C`);
    });
    
    // Listen to custom signal
    heater.on('stateChanged', function(state) {
        console.log(`Heater state signal: ${state ? "ON" : "OFF"}`);
    });

    // Initial setup
    thermostat.setMode('auto');
    thermostat.setTargetTemperature(23.0); // Want it a bit warmer

    // Simulate temperature changes over time
    const simulationSteps = 10;
    for (let i = 0; i < simulationSteps; i++) {
        tempSensor.update();
    }

    return {
        finalTemperature: tempSensor.currentTemperature,
        heaterState: heater.isOn,
        targetTemperature: thermostat.targetTemperature
    };
}
