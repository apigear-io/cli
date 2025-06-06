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

// Heater methods
heater.turnOn = function () {
    if (!heater.isOn) {
        heater.isOn = true;
        console.log("Heater turned ON");
    }
}

heater.turnOff = function () {
    if (heater.isOn) {
        heater.isOn = false;
        console.log("Heater turned OFF");
    }
}

heater.updateTemperature = function (deltaTime) {
    if (heater.isOn) {
        // Simple temperature increase model
        // Temperature rises faster when difference to max temp is larger
        const heatIncrease = (heater.maxTemp - heater.temperature) * 0.1;
        heater.temperature += heatIncrease * (deltaTime / 1000);
    } else {
        // Natural cooling model
        // Temperature falls faster when difference to ambient temp is larger
        const cooling = (heater.temperature - tempSensor.currentTemperature) * 0.05;
        heater.temperature -= cooling * (deltaTime / 1000);
    }
}

// Thermostat methods
thermostat.setTargetTemperature = function (temp) {
    if (temp >= heater.minTemp && temp <= heater.maxTemp) {
        thermostat.targetTemperature = temp;
        console.log(`Target temperature set to ${temp}°C`);
        thermostat.checkTemperature();
    } else {
        console.log(`Temperature ${temp}°C is outside allowed range`);
    }
}

thermostat.checkTemperature = function () {
    const currentTemp = tempSensor.currentTemperature;
    const lowerBound = thermostat.targetTemperature - thermostat.tolerance;
    const upperBound = thermostat.targetTemperature + thermostat.tolerance;

    if (currentTemp < lowerBound) {
        heater.turnOn();
    } else if (currentTemp > upperBound) {
        heater.turnOff();
    }
}

thermostat.setMode = function (newMode) {
    if (newMode === 'auto' || newMode === 'manual') {
        thermostat.mode = newMode;
        console.log(`Thermostat mode set to ${newMode}`);
        if (newMode === 'auto') {
            thermostat.checkTemperature();
        }
    }
}

// Temperature sensor methods
tempSensor.update = function () {
    const now = Date.now();
    const deltaTime = now - tempSensor.lastUpdate;
    tempSensor.lastUpdate = now;

    // Update current temperature based on heater's influence
    const heatTransfer = (heater.temperature - tempSensor.currentTemperature) * 0.1;
    tempSensor.currentTemperature += heatTransfer * (deltaTime / 1000);

    // Add some random fluctuation
    tempSensor.currentTemperature += (Math.random() - 0.5) * 0.1;

    console.log(`Current temperature: ${tempSensor.currentTemperature.toFixed(1)}°C`);

    if (thermostat.mode === 'auto') {
        thermostat.checkTemperature();
    }
}

function main() {
    // Set up monitoring
    heater.$.onProperty("isOn", function (isOn) {
        console.log(`Heater state changed to: ${isOn ? "ON" : "OFF"}`);
    });

    tempSensor.$.onProperty("currentTemperature", function (temp) {
        console.log(`Temperature sensor reading: ${temp.toFixed(1)}°C`);
    });

    thermostat.$.onProperty("targetTemperature", function (temp) {
        console.log(`Target temperature changed to: ${temp.toFixed(1)}°C`);
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
