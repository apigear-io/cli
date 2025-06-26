// create a simulation service with natural API
const counter = $createService("demo.Counter", {
    count: 0,
    max: 10
});

// Natural event handling - no need for service.$
counter.on("count", function (count) {
    console.log("count:", count);
});

// Natural method definition with auto-bound 'this'
counter.increment = function() {
    console.log("incremented from", this.count);
    if (this.count >= this.max) {
        this.reset();
    } else {
        this.count = this.count + 1;  // Natural property assignment
    }
    console.log("incremented to", this.count);
}

counter.reset = function() {
    const oldCount = this.count;
    this.count = 0;
    // Streamlined signal emission
    this.emit("reset", oldCount, this.count);
}

// Add a decrement method for completeness
counter.decrement = function() {
    if (this.count > 0) {
        this.count--;
    }
}

function main() {
    console.log("=== Counter Demo ===");
    console.log("Initial count:", counter.count);
    
    // Demonstrate natural API usage
    counter.increment();
    counter.increment();
    counter.increment();
    
    console.log("After 3 increments:", counter.count);
    
    // Test reset functionality
    counter.count = 9;  // Direct property assignment
    console.log("Set count to 9:", counter.count);
    
    counter.increment();  // Should trigger reset
    console.log("Final count:", counter.count);
    
    // Exit after demo
    console.log("Demo completed!");
    if (typeof $quit === 'function') {
        setTimeout($quit, 1000);  // Exit after 1 second
    }
}