# vfs

Virtual embedded file system for demo templates.

## Purpose

The `vfs` package provides embedded demo/template files that are compiled directly into the Go binary. These files serve as boilerplate templates for creating new APIGear projects.

## Key Exports

All exports are `[]byte` variables containing embedded file contents:

- `DemoModuleYaml` - YAML template for module configuration
- `DemoSolutionYaml` - YAML template for solution configuration
- `DemoModuleIdl` - IDL template for module definitions
- `DemoSimulationJs` - JavaScript template for simulation logic

## Usage

These templates are used by the `prj` package when initializing new projects with demo content.

## Dependencies

This package has no dependencies on other `pkg/` packages.
