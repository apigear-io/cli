const state = $createService("vehicle.State", {
    location: { x: 0, y: 0 },
    speed: 0,
    rpm: 0,
    fuelLevel: 0,
    fuelLevelWarning: false,
    temperature: 0,
    overheatWarning: false
});

const indicators = $createService("vehicle.Indicators", {
    checkEngine: false,
    oilPressure: false,
    battery: false,
    airbag: false,
    brake: false,
    seatbelt: false,
    tractionControl: false,
    highBeam: false
});

const commands = $createService("vehicle.Commands", {});

// Command methods using natural API
commands.turnOn = function () {
    const order = ['checkEngine', 'oilPressure', 'battery', 'brake', 'seatbelt', 'tractionControl', 'highBeam'];
    let index = 0;
    const interval = setInterval(function() {
        if (index < order.length) {
            const indicator = order[index];
            indicators[indicator] = true;
            console.log(`Turned on ${indicator}`);
            index++;
        } else {
            clearInterval(interval);
            commands.emit('allIndicatorsOn');
        }
    }, 200);
}

commands.turnOff = function () {
    indicators.checkEngine = false;
    indicators.oilPressure = false;
    indicators.battery = false;
    indicators.airbag = false;
    indicators.brake = false;
    indicators.seatbelt = false;
    indicators.tractionControl = false;
    indicators.highBeam = false;
    this.emit('allIndicatorsOff');
}

// Monitor indicators using natural API
indicators.on("checkEngine", function (value) {
    console.log("checkEngine changed:", value);
});

// Add method to state service for speed updates
state.accelerate = function(amount = 10) {
    this.speed += amount;
    this.rpm = Math.min(8000, this.speed * 100);
    
    // Update fuel consumption
    this.fuelLevel = Math.max(0, this.fuelLevel - amount * 0.01);
    this.fuelLevelWarning = this.fuelLevel < 10;
    
    // Update temperature
    this.temperature = Math.min(120, this.temperature + amount * 0.1);
    this.overheatWarning = this.temperature > 100;
}

function main() {
    // Set up event monitoring
    state.on('speed', function(speed) {
        console.log(`Speed: ${speed} km/h`);
    });
    
    state.on('fuelLevelWarning', function(warning) {
        if (warning) {
            console.log('⚠️  Low fuel warning!');
        }
    });
    
    state.on('overheatWarning', function(warning) {
        if (warning) {
            console.log('⚠️  Engine overheating!');
        }
    });
    
    commands.on('allIndicatorsOn', function() {
        console.log('All indicators checked');
    });
    
    // Initialize
    state.fuelLevel = 50; // Start with 50% fuel
    state.temperature = 20; // Cold engine
    
    // Run startup sequence
    commands.turnOn();
    
    // Simulate driving
    let drivingInterval = setInterval(function() {
        state.accelerate();
        if (state.speed >= 120 || state.fuelLevel <= 0) {
            clearInterval(drivingInterval);
            console.log('Stopping simulation');
            commands.turnOff();
        }
    }, 500);
}