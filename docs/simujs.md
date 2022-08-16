# Simulation

- Several playbooks per simulation
- No sequence, just a number of steps
- playbooks can have names and run via simulation ui
- ui is done as terminal ui or web ui
- ui can be used to run playbooks
- playbooks can also be independent of an interface
- animations must have easing curves or patterns
- animations can be done in parallel or in sequence
- you can link an animation to a API modules to validate correctness and types
- simulation can also be done using fake values

## Operation Behavior

Behavior for operations can be attached using JS script. The script is called every time the operation is called.

```javascript
function add(a, b) {
  return a + b;
}
```

The properties of the object are accessible as `$props` and the parameters are passed into the function.

```javascript
function increment(value) {
  $props.count += value;
}
```

Property change signals are automatically sent back to the calling client or can be manually triggered.

```javascript
function increment(value) {
  $change("count", $props.count + value);
}
```

As part of a operation call, it is also possible to emit signals.

```javascript
function increment(value) {
  // signal system shutdown in 5 secs
  $emit("shutdown", 5);
}
```

## Playbooks

It is possible to define playbooks which can be run from the UI or triggered automatically on simulation start.

A simulation scenario can have one or more playbooks attached which can be run in sequence or in parallel and using different timings.

```javascript
let counter = $.get("demo.Counter");
for (let i = 0; i < 10; i++) {
  counter.call("increment", 1);
  counter.set("count", counter.get("count") + 1);
  wait(1000);
}
```
