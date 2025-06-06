package vfs

import _ "embed"

//go:embed demo.module.yaml
var DemoModuleYaml []byte

//go:embed demo.solution.yaml
var DemoSolutionYaml []byte

//go:embed demo.module.idl
var DemoModuleIdl []byte

//go:embed demo.sim.js
var DemoSimulationJs []byte
