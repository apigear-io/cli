package vfs

import _ "embed"

//go:embed demo.module.yaml
var DemoModuleYaml []byte

//go:embed demo.solution.yaml
var DemoSolutionYaml []byte

//go:embed demo.scenario.yaml
var DemoScenarioYaml []byte

//go:embed demo.module.idl
var DemoModuleIdl []byte

//go:embed demo.service.js
var DemoServiceJs []byte

//go:embed demo.client.js
var DemoClientJs []byte
