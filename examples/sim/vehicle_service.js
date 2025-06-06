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

commands.turnOn = function () {
    const order = [checkEngine, oilPressure, battery, brake, seatbelt, tractionControl, highBeam]
    setInterval(function() {
        for(const entry in order) {
            indicators[entry] = true
        }
    }, 200)
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
}

indicators.$.onProperty("checkEngine", function (value) {
    console.log("checkEngine changed", value);
});

function main() {
    commands.turnOn();
    setInterval(function() {
        state.speed+= 10
    }, 200)
}