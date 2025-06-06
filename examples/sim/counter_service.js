// a bare service is a service which uses the bare methods
const service = $createService("counter", { count: 1 });

service.increment = function () {
  console.log("called service increment");
  service.count++
};

service.$.onProperty("count", function (value) {
  console.log("on property service count changed", value);
});

function main() {
  service.increment();
  return service.count;
}
