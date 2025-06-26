// Counter service example using the natural API
const counter = $createService("counter", { count: 1 });

// Define methods using natural function assignment
counter.increment = function () {
  console.log("called counter increment");
  this.count++;  // Natural property access with 'this'
};

// Use the streamlined event handler for property changes
counter.on("count", function (value) {
  console.log("counter.count changed to:", value);
});

function main() {
  console.log("Initial count:", counter.count);  // Natural property read
  counter.increment();
  console.log("Final count:", counter.count);
  return counter.count;
}
