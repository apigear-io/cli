// Counter service using bare service API (without proxy)
// This shows the underlying API that the proxy wraps
const service = $createBareService("counter", { count: 1 });

service.onMethod("increment", function () {
  console.log("called service increment");
  const count = service.getProperty("count");
  service.setProperty("count", count + 1);
});

service.onProperty("count", function (value) {
  console.log("on property service count changed", value);
});

// Note: The natural API with proxy would be:
// const counter = $createService("counter", { count: 1 });
// counter.increment = function() { this.count++; };
// counter.on("count", function(value) { ... });

