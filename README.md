# go-template

A lightweight CLI tool for rendering templates using Go's text/template engine with Sprig functions.

## Features

- Process templates from files or standard input
- Inject parameters from YAML files
- Full support for [Sprig](https://masterminds.github.io/sprig/) template functions
- Simple, focused interface

## Installation

```sh
go install github.com/conao3/go-template@latest
```

Or build from source:

```sh
git clone https://github.com/conao3/go-template.git
cd go-template
go install .
```

## Usage

```sh
go-template -i <template-file> [-p <params-file>]
```

### Options

| Flag | Description |
|------|-------------|
| `-i, --input` | Input template file (use `-` for stdin) |
| `-p, --param` | YAML file containing template parameters |

### Examples

Basic template rendering:

```sh
$ go-template -i sample/01_simple.yml
- a
- b
- c
```

Using range to iterate:

```sh
$ go-template -i sample/02_range.yml
- 0
- 1
- 2
- 3
- 4
```

Injecting values from a parameter file:

```sh
$ go-template -i sample/03_injection.yml -p sample/03_injection.val.yml
[1 2 3]

$ go-template -i sample/03_injection_list.yml -p sample/03_injection.val.yml
- item_1
- item_2
- item_3
```

## License

See [LICENSE](LICENSE) for details.
