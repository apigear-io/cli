// Edge case tests for the proxy implementation
// Tests unusual scenarios and boundary conditions

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
// Edge Case Tests
// ====================

test("Method overwriting property", () => {
    const service = $createService("test.Overwrite", {
        value: 10
    });
    
    // Overwrite property with method
    service.value = function() {
        return 42;
    };
    
    assert(typeof service.value === 'function', "value should be a function");
    assertEqual(service.value(), 42, "method return value");
});

test("Property overwriting method", () => {
    const service = $createService("test.OverwriteMethod", {});
    
    service.method = function() {
        return "original";
    };
    
    assertEqual(service.method(), "original", "original method");
    
    // Overwrite method with property
    service.method = "not a function";
    assertEqual(service.method, "not a function", "overwritten with property");
});

test("Recursive method calls", () => {
    const service = $createService("test.Recursive", {
        depth: 0
    });
    
    service.recurse = function(n) {
        if (n <= 0) return this.depth;
        this.depth++;
        return this.recurse(n - 1);
    };
    
    const result = service.recurse(5);
    assertEqual(result, 5, "recursive call result");
    assertEqual(service.depth, 5, "depth after recursion");
});

test("Method with no return value", () => {
    const service = $createService("test.NoReturn", {
        sideEffect: false
    });
    
    service.doSomething = function() {
        this.sideEffect = true;
        // No explicit return
    };
    
    const result = service.doSomething();
    assertEqual(result, undefined, "no return value");
    assert(service.sideEffect, "side effect occurred");
});

test("Special property names", () => {
    const service = $createService("test.SpecialNames", {
        "constructor": "not a constructor",
        "prototype": "not a prototype",
        "__proto__": "not proto",
        "toString": "not toString"
    });
    
    assertEqual(service.constructor, "not a constructor", "constructor property");
    assertEqual(service.prototype, "not a prototype", "prototype property");
    assertEqual(service.__proto__, "not proto", "__proto__ property");
    assertEqual(service.toString, "not toString", "toString property");
});

test("Properties starting with underscore", () => {
    const service = $createService("test.Underscore", {
        _private: "private value",
        public: "public value"
    });
    
    assertEqual(service._private, "private value", "underscore property");
    assertEqual(service.public, "public value", "public property");
    
    // Set new underscore property
    service._newPrivate = "new private";
    assertEqual(service._newPrivate, "new private", "new underscore property");
});

test("Method throwing exception", () => {
    const service = $createService("test.Exception", {});
    
    service.throwError = function() {
        throw new Error("Method error");
    };
    
    let caught = false;
    try {
        service.throwError();
    } catch (e) {
        caught = true;
        assert(e.message === "Method error", "error message");
    }
    assert(caught, "exception was caught");
});

test("Circular reference in properties", () => {
    const service = $createService("test.Circular", {
        name: "service"
    });
    
    // Create circular reference
    service.self = service;
    
    assert(service.self === service, "circular reference");
    assertEqual(service.self.name, "service", "access through circular ref");
    assertEqual(service.self.self.self.name, "service", "deep circular access");
});

test("Large number of properties", () => {
    const props = {};
    for (let i = 0; i < 1000; i++) {
        props[`prop${i}`] = i;
    }
    
    const service = $createService("test.ManyProps", props);
    
    assertEqual(service.prop0, 0, "first property");
    assertEqual(service.prop500, 500, "middle property");
    assertEqual(service.prop999, 999, "last property");
    
    // Check enumeration works
    const keys = Object.keys(service);
    assert(keys.length >= 1000, "all properties enumerable");
});

test("Empty service", () => {
    const service = $createService("test.Empty", {});
    
    // Should still have proxy methods
    assert('on' in service, "has on method");
    assert('emit' in service, "has emit method");
    assert('$' in service, "has $ property");
    
    // Can add properties
    service.newProp = "value";
    assertEqual(service.newProp, "value", "can add property");
});

test("Null and undefined values", () => {
    const service = $createService("test.NullUndefined", {
        nullProp: null,
        undefinedProp: undefined
    });
    
    assertEqual(service.nullProp, null, "null property");
    assertEqual(service.undefinedProp, undefined, "undefined property");
    
    // Set to null/undefined
    service.newNull = null;
    service.newUndefined = undefined;
    
    assertEqual(service.newNull, null, "new null property");
    assertEqual(service.newUndefined, undefined, "new undefined property");
});

test("Method modifying other properties", () => {
    const service = $createService("test.CrossModify", {
        a: 1,
        b: 2,
        c: 3
    });
    
    service.shuffle = function() {
        const temp = this.a;
        this.a = this.b;
        this.b = this.c;
        this.c = temp;
    };
    
    service.shuffle();
    assertEqual(service.a, 2, "a after shuffle");
    assertEqual(service.b, 3, "b after shuffle");
    assertEqual(service.c, 1, "c after shuffle");
});

test("Property listeners with same name as methods", () => {
    const service = $createService("test.NameConflict", {
        value: 0
    });
    
    // Add a method named 'value'
    service.getValue = function() {
        return this.value;
    };
    
    let listenerCalled = false;
    // Listen to property 'value', not method 'getValue'
    service.on('value', function() {
        listenerCalled = true;
    });
    
    service.value = 10;
    assert(listenerCalled, "property listener called");
    assertEqual(service.getValue(), 10, "method still works");
});

test("Methods with 'arguments' object", () => {
    const service = $createService("test.Arguments", {});
    
    service.sum = function() {
        let total = 0;
        for (let i = 0; i < arguments.length; i++) {
            total += arguments[i];
        }
        return total;
    };
    
    assertEqual(service.sum(1, 2, 3), 6, "sum with 3 args");
    assertEqual(service.sum(1, 2, 3, 4, 5), 15, "sum with 5 args");
    assertEqual(service.sum(), 0, "sum with no args");
});

test("Property change during listener", () => {
    const service = $createService("test.ChangeInListener", {
        value: 0,
        other: 0
    });
    
    service.on('value', function(newValue) {
        // Change another property during listener
        service.other = newValue * 2;
    });
    
    service.value = 5;
    assertEqual(service.other, 10, "other property changed in listener");
});

test("Multiple signals with same listener", () => {
    const service = $createService("test.MultiSignal", {});
    
    let eventCount = 0;
    let lastEvent = null;
    
    const handler = function(event) {
        eventCount++;
        lastEvent = event;
    };
    
    service.on('event1', handler);
    service.on('event2', handler);
    
    service.emit('event1', 'first');
    assertEqual(eventCount, 1, "first event");
    assertEqual(lastEvent, 'first', "first event data");
    
    service.emit('event2', 'second');
    assertEqual(eventCount, 2, "second event");
    assertEqual(lastEvent, 'second', "second event data");
});

test("Method returning another method", () => {
    const service = $createService("test.MethodFactory", {
        multiplier: 2
    });
    
    service.createMultiplier = function(factor) {
        const self = this;
        return function(value) {
            return value * factor * self.multiplier;
        };
    };
    
    const times3 = service.createMultiplier(3);
    assertEqual(times3(5), 30, "factory method result");
});

test("Boolean property coercion", () => {
    const service = $createService("test.BoolCoercion", {
        flag: false
    });
    
    service.toggle = function() {
        this.flag = !this.flag;
        return this.flag;
    };
    
    assert(service.toggle() === true, "first toggle");
    assert(service.toggle() === false, "second toggle");
    assert(service.toggle() === true, "third toggle");
});

test("String concatenation in methods", () => {
    const service = $createService("test.StringConcat", {
        prefix: "Hello",
        suffix: "World"
    });
    
    service.join = function(separator) {
        return this.prefix + separator + this.suffix;
    };
    
    assertEqual(service.join(" "), "Hello World", "space separator");
    assertEqual(service.join("-"), "Hello-World", "dash separator");
    assertEqual(service.join(""), "HelloWorld", "no separator");
});

// ====================
// Test Summary
// ====================

console.log("\n" + "=".repeat(50));
console.log(`Edge case tests passed: ${testsPassed}`);
console.log(`Edge case tests failed: ${testsFailed}`);
console.log("=".repeat(50));

if (testsFailed > 0) {
    console.log("\nFailed tests:");
    errors.forEach(error => console.log(`  - ${error}`));
    throw new Error(`${testsFailed} edge case tests failed`);
}

// Return success indicator
"ALL_EDGE_TESTS_PASSED";