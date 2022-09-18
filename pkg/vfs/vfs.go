package vfs

import _ "embed"

//go:embed demo.module.yaml
var DemoModule []byte

//go:embed demo.solution.yaml
var DemoSolution []byte

//go:embed demo.scenario.yaml
var DemoScenario []byte
