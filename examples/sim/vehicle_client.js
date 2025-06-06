const channel = $createChannel();
const commands = channel.createClient("vehicle.Commands");

function main() {
    commands.callMethod("turnOff");
}