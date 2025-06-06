// service side
const service = $createBareService("counter", { count: 1 });

service.onMethod("increment", function () {
  console.log("called service increment");
  const count = service.getProperty("count");
  service.setProperty("count", count + 1);
});

service.onProperty("count", function (value) {
  console.log("on property service count changed", value);
});

