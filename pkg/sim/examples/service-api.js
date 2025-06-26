// Service API Examples - Natural Usage Patterns

// ============= Basic Service Creation =============
const counter = $createService("demo.Counter", {
    count: 0,
    max: 10
});

// ============= Property Access =============
// Reading properties - natural access
console.log(counter.count);  // 0
console.log(counter.max);    // 10

// Writing properties - natural assignment
counter.count = 5;
counter.max = 100;

// ============= Method Definition =============
// Define methods using natural function assignment
counter.increment = function() {
    // 'this' is automatically bound to the service proxy
    if (this.count < this.max) {
        this.count++;
    }
};

counter.decrement = function() {
    if (this.count > 0) {
        this.count--;
    }
};

counter.reset = function() {
    this.count = 0;
    this.emit('reset');  // Emit signal
};

// ============= Event Handling =============
// Listen to property changes
counter.on('count', function(newValue) {
    console.log('Count changed to:', newValue);
});

// Listen to signals
counter.on('reset', function() {
    console.log('Counter was reset!');
});

// ============= Signal Emission =============
// Emit custom signals with arguments
counter.emit('custom', 'arg1', 'arg2');

// ============= Advanced Patterns =============

// 1. Method Chaining Pattern
counter.setCount = function(value) {
    this.count = value;
    return this;  // Enable chaining
};

counter.setMax = function(value) {
    this.max = value;
    return this;  // Enable chaining
};

// Usage: counter.setCount(5).setMax(20);

// 2. Computed Properties Pattern
counter.percentage = function() {
    return (this.count / this.max) * 100;
};

// 3. Validation Pattern
counter.safeIncrement = function(amount = 1) {
    const newCount = this.count + amount;
    if (newCount <= this.max && newCount >= 0) {
        this.count = newCount;
        return true;
    }
    return false;
};

// 4. Async Operations Pattern
counter.delayedReset = function(delay) {
    const self = this;
    setTimeout(function() {
        self.reset();
    }, delay);
};

// ============= Access to Raw Service =============
// Use counter.$ to access the underlying service object
// This is useful for advanced operations
const rawService = counter.$;
rawService.onProperty('count', function(value) {
    // Direct property listener
});

// ============= Error Handling =============
// The proxy provides helpful warnings for undefined properties
// console.log(counter.nonExistent);  // Warning: Property 'nonExistent' not found

// ============= Best Practices =============
// 1. Use natural property access (counter.count) instead of getProperty/setProperty
// 2. Define methods with regular function assignment
// 3. Use 'on' for both property changes and signals
// 4. Use 'emit' for signal emission
// 5. Access raw service with .$ only when necessary
// 6. Leverage 'this' binding in methods for cleaner code