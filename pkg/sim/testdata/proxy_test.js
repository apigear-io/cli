// Test suite for the Go-based proxy implementation
// This ensures the proxy works correctly from JavaScript

let testsPassed = 0;
let testsFailed = 0;
const errors = [];

function assert(condition, message) {
    if (!condition) {
        const error = `Assertion failed: ${message}`;
        errors.push(error);
        throw new Error(error);
    }
}

function assertEqual(actual, expected, message) {
    if (actual !== expected) {
        const error = `${message}: expected ${JSON.stringify(expected)}, got ${JSON.stringify(actual)}`;
        errors.push(error);
        throw new Error(error);
    }
}

function test(name, testFn) {
    try {
        testFn();
        console.log(`✓ ${name}`);
        testsPassed++;
    } catch (e) {
        console.log(`✗ ${name}: ${e.message}`);
        testsFailed++;
    }
}

// ====================
// Test Suite
// ====================

test("Basic property access", () => {
    const service = $createService("test.Basic", {
        name: "test",
        count: 42,
        active: true
    });
    
    assertEqual(service.name, "test", "name property");
    assertEqual(service.count, 42, "count property");
    assertEqual(service.active, true, "active property");
});

test("Property modification", () => {
    const service = $createService("test.Modify", {
        value: 10
    });
    
    service.value = 20;
    assertEqual(service.value, 20, "modified value");
    
    service.newProp = "dynamic";
    assertEqual(service.newProp, "dynamic", "dynamically added property");
});

test("Method definition and invocation", () => {
    const service = $createService("test.Methods", {
        counter: 0
    });
    
    service.increment = function() {
        this.counter++;
        return this.counter;
    };
    
    const result = service.increment();
    assertEqual(result, 1, "increment return value");
    assertEqual(service.counter, 1, "counter after increment");
});

test("'this' binding in methods", () => {
    const service = $createService("test.ThisBinding", {
        value: 5,
        multiplier: 2
    });
    
    service.calculate = function() {
        // 'this' should be bound to the proxy
        return this.value * this.multiplier;
    };
    
    assertEqual(service.calculate(), 10, "method using 'this'");
    
    // Test that 'this' is not undefined
    service.checkThis = function() {
        assert(this !== undefined, "'this' is undefined");
        assert(this !== null, "'this' is null");
        return true;
    };
    
    assert(service.checkThis(), "'this' check failed");
});

test("Method chaining", () => {
    const service = $createService("test.Chaining", {
        value: 0
    });
    
    service.add = function(n) {
        this.value += n;
        return this;  // Return 'this' for chaining
    };
    
    service.multiply = function(n) {
        this.value *= n;
        return this;
    };
    
    service.add(5).multiply(3).add(2);
    assertEqual(service.value, 17, "chained operations");
});

test("Methods calling other methods", () => {
    const service = $createService("test.MethodCalls", {
        count: 0
    });
    
    service.increment = function() {
        this.count++;
    };
    
    service.incrementTwice = function() {
        this.increment();
        this.increment();
    };
    
    service.incrementTwice();
    assertEqual(service.count, 2, "method calling another method");
});

test("Property change notifications", () => {
    const service = $createService("test.PropertyNotify", {
        value: 0
    });
    
    let notificationCount = 0;
    let lastValue = null;
    
    service.on('value', function(newValue) {
        notificationCount++;
        lastValue = newValue;
    });
    
    service.value = 10;
    assertEqual(notificationCount, 1, "first notification");
    assertEqual(lastValue, 10, "first value");
    
    service.value = 20;
    assertEqual(notificationCount, 2, "second notification");
    assertEqual(lastValue, 20, "second value");
    
    // Same value should not trigger notification
    service.value = 20;
    assertEqual(notificationCount, 2, "no notification for same value");
});

test("Signal emission and handling", () => {
    const service = $createService("test.Signals", {});
    
    let signalReceived = false;
    let receivedArgs = null;
    
    service.on('customSignal', function(arg1, arg2, arg3) {
        signalReceived = true;
        receivedArgs = [arg1, arg2, arg3];
    });
    
    service.emit('customSignal', 'a', 'b', 'c');
    
    assert(signalReceived, "signal not received");
    assertEqual(receivedArgs[0], 'a', "first arg");
    assertEqual(receivedArgs[1], 'b', "second arg");
    assertEqual(receivedArgs[2], 'c', "third arg");
});

test("Multiple listeners", () => {
    const service = $createService("test.MultipleListeners", {
        value: 0
    });
    
    let count1 = 0;
    let count2 = 0;
    
    service.on('value', function() {
        count1++;
    });
    
    service.on('value', function() {
        count2++;
    });
    
    service.value = 10;
    assertEqual(count1, 1, "first listener");
    assertEqual(count2, 1, "second listener");
});

test("Raw service access via $", () => {
    const service = $createService("test.RawAccess", {
        prop: "value"
    });
    
    const raw = service.$;
    assert(raw !== undefined, "raw service is undefined");
    assert(raw !== null, "raw service is null");
    assert(typeof raw === 'object', "raw service is not an object");
});

test("Property enumeration", () => {
    const service = $createService("test.Enumeration", {
        prop1: "value1",
        prop2: "value2"
    });
    
    service.method1 = function() { return "result"; };
    
    const keys = Object.keys(service);
    assert(keys.includes('prop1'), "prop1 not enumerable");
    assert(keys.includes('prop2'), "prop2 not enumerable");
    assert(keys.includes('method1'), "method1 not enumerable");
    assert(keys.includes('on'), "on not enumerable");
    assert(keys.includes('emit'), "emit not enumerable");
    assert(keys.includes('$'), "$ not enumerable");
});

test("Property existence check", () => {
    const service = $createService("test.PropertyExistence", {
        exists: true
    });
    
    service.method = function() {};
    
    assert('exists' in service, "property not found");
    assert('method' in service, "method not found");
    assert('on' in service, "on not found");
    assert('emit' in service, "emit not found");
    assert('$' in service, "$ not found");
    assert(!('nonExistent' in service), "non-existent property found");
});

test("Complex scenario with calculator", () => {
    const calc = $createService("test.Calculator", {
        result: 0,
        operations: []
    });
    
    calc.add = function(n) {
        this.result += n;
        this.operations.push(`add ${n}`);
        return this;
    };
    
    calc.subtract = function(n) {
        this.result -= n;
        this.operations.push(`subtract ${n}`);
        return this;
    };
    
    calc.multiply = function(n) {
        this.result *= n;
        this.operations.push(`multiply ${n}`);
        return this;
    };
    
    calc.clear = function() {
        this.result = 0;
        this.operations = [];
        this.emit('cleared');
        return this;
    };
    
    let clearedCount = 0;
    calc.on('cleared', function() {
        clearedCount++;
    });
    
    calc.add(10).multiply(2).subtract(5);
    assertEqual(calc.result, 15, "calculation result");
    assertEqual(calc.operations.length, 3, "operations count");
    
    calc.clear();
    assertEqual(calc.result, 0, "cleared result");
    assertEqual(calc.operations.length, 0, "cleared operations");
    assertEqual(clearedCount, 1, "clear event emitted");
});

test("Method with arguments and return value", () => {
    const service = $createService("test.MethodArgs", {});
    
    service.greet = function(name, title) {
        return `Hello ${title} ${name}!`;
    };
    
    const greeting = service.greet("Smith", "Mr.");
    assertEqual(greeting, "Hello Mr. Smith!", "method with args");
});

test("Method stored in variable", () => {
    const service = $createService("test.MethodVariable", {
        value: 10
    });
    
    service.getValue = function() {
        return this.value;
    };
    
    const method = service.getValue;
    // Calling through variable should still have 'this' bound
    // when called with the service as context
    const result = method.call(service);
    assertEqual(result, 10, "method called through variable");
});

test("Nested property access", () => {
    const service = $createService("test.Nested", {
        config: {
            host: "localhost",
            port: 8080
        }
    });
    
    assert(service.config !== undefined, "config is undefined");
    assertEqual(service.config.host, "localhost", "nested host");
    assertEqual(service.config.port, 8080, "nested port");
    
    // Modify nested property
    service.config = { host: "example.com", port: 3000 };
    assertEqual(service.config.host, "example.com", "modified nested host");
});

test("Array property handling", () => {
    const service = $createService("test.Arrays", {
        items: [1, 2, 3]
    });
    
    assertEqual(service.items.length, 3, "array length");
    assertEqual(service.items[0], 1, "first element");
    
    // Modify array
    service.items = [4, 5, 6, 7];
    assertEqual(service.items.length, 4, "modified array length");
    assertEqual(service.items[2], 6, "third element of modified array");
});

test("Undefined property access", () => {
    const service = $createService("test.Undefined", {
        defined: "value"
    });
    
    assertEqual(service.undefined, undefined, "undefined property");
    assertEqual(service.nonExistent, undefined, "non-existent property");
});

test("Property types preservation", () => {
    const service = $createService("test.Types", {
        string: "text",
        number: 42,
        boolean: true,
        null: null,
        array: [1, 2, 3],
        object: { key: "value" }
    });
    
    assertEqual(typeof service.string, "string", "string type");
    assertEqual(typeof service.number, "number", "number type");
    assertEqual(typeof service.boolean, "boolean", "boolean type");
    assertEqual(service.null, null, "null value");
    assert(Array.isArray(service.array), "array type");
    assertEqual(typeof service.object, "object", "object type");
});

// ====================
// Test Summary
// ====================

console.log("\n" + "=".repeat(50));
console.log(`Tests passed: ${testsPassed}`);
console.log(`Tests failed: ${testsFailed}`);
console.log("=".repeat(50));

if (testsFailed > 0) {
    console.log("\nFailed tests:");
    errors.forEach(error => console.log(`  - ${error}`));
    throw new Error(`${testsFailed} tests failed`);
}

// Return success indicator
"ALL_TESTS_PASSED";