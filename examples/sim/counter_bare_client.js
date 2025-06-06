// client side
const channel = $createChannel();
const client = channel.createClient("counter");

client.onProperty("count", function (value) {
  console.log("client count changed", value);
});

function main() {
  console.log("main");
  for (let i = 0; i < 1; i++) {
    console.log("increment");
    client.callMethod("increment");
  }
}
