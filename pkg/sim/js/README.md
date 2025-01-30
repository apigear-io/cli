# Simulation JS Runtime

This package provides the JavaScript execution environment for the simulation engine. It consist of Goja, a JavaScript runtime written in Go.

The js specific types are defined in the `pkg/js` package and functions are prefixed with `js.`. Other functions are internal. Public APIs use type go types.

## Types

A `jsRegister` function per type is provided to register the type with the runtime as js object. This is used in the `pkg/js/world.go` and `pkg/js/actor.go` files.

For global functions (such as constructors), a `js{type}Register` function is provided in the `pkg/js/world.go` file and run when a world is created.

## World and Actors

The world is the main container for the simulation. It contains actors and hooks. Actors are the main entities in the simulation. They contain state, methods and signals. Hooks are used to hook into the simulation lifecycle.