# go-template

Generic tempalte engine powered by Golang text/template.

## Install

```sh
go install .
```

## Usage

```sh
$ go-template -i sample/01_simple.yml
- a
- b
- c

$ go-template -i sample/02_range.yml
- 0
- 1
- 2
- 3
- 4

$ go-template -i sample/03_injection.yml -p sample/03_injection.val.yml
[1 2 3]

$ go-template -i sample/03_injection_list.yml -p sample/03_injection.val.yml
- item_1
- item_2
- item_3
```
