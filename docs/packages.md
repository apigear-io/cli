# Template Package

Packages are created using the CLI tool. A package is a templates folder with a README.md, rules.yaml and a manifest.yaml file.

- README.md: - The README.md file is the main documentation for the package.
- manifest.yaml - The manifest.yaml file is used to define the package, author and license.
- rules.yaml - The rules.yaml file is used to define the rules for the package.
- templates/ - The templates folder is where the templates are stored.

The CLI can create new packages, update existing packages, and delete packages as also publish packages to the registry.

## Package Server

The package server allows a user to upload packages to the registry and other users to download packages from the registry.

## Usage

TBD: Create a new package:

```sh
$ cli pkg init <package-name>
```

TBD: Publish a package:

```sh
$ cli pkg publish <package-name>
```

TBD: Pack a package:

```sh
$ cli pkg pack <package-name>
```

TBD: Search for packages:

```sh
$ cli pkg search <package-name>
```

TBD: List installed packages

```
$ cli pkg list <package-name>
```

TBD: Remove a package:

```sh
$ cli pkg remove <package-name>
```

TBD: Detailed information about a package:

```sh
$ cli pkg info <package-name>
```
