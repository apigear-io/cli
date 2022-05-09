# Code Generation

Code generation is based on an IDL format and a JSON (YAML) Schema format as input. The code generator uses a template package to create the code in the output folder.

## Usage

TODO: Generate code from a solution document

```sh
$ cli codegen --sol <solution-document>
```

TODO: Generate code using the expert mode

```sh
$ cli codegen --input apis --output ./out --template ./templates
```

- `--watch`: It is possible watch over the input and templates files and regenerate the code when they are changed.
- `--force`: It is possible to force the regeneration of the code.
