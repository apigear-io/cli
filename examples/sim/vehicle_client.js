// Vehicle client - connects to remote vehicle services
const channel = $createChannel();
const commands = channel.createClient("vehicle.Commands");
const state = channel.createClient("vehicle.State");
const indicators = channel.createClient("vehicle.Indicators");

// Monitor state changes
state.onProperty("speed", function(speed) {
    console.log(`Client - Speed: ${speed} km/h`);
});

state.onProperty("fuelLevelWarning", function(warning) {
    if (warning) {
        console.log("Client - Low fuel warning!");
    }
});

indicators.onProperty("checkEngine", function(value) {
    console.log(`Client - Check engine: ${value}`);
});

function main() {
    console.log("Vehicle client starting...");
    
    // Turn on vehicle systems
    commands.callMethod("turnOn");
    
    // Wait a bit then turn off
    setTimeout(function() {
        console.log("Turning off vehicle systems...");
        commands.callMethod("turnOff");
    }, 3000);
}