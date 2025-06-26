// Counter client - connects to a remote counter service
// Note: Channel clients communicate remotely and don't use the proxy API
const channel = $createChannel();
const client = channel.createClient("counter");

client.onProperty("count", function (value) {
  console.log("client: count changed to", value);
});

function main() {
  console.log("Counter client started");
  
  // Call the remote increment method multiple times
  for (let i = 0; i < 5; i++) {
    console.log(`Calling increment (${i + 1}/5)`);
    client.callMethod("increment");
  }
  
  console.log("All increment calls sent");
}
