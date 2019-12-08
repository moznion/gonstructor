# gonstructor

A command-line tool to generate a constructor for the struct.

## Usage

```
Usage of gonstructor:
  -constructorTypes string
        [optional] comma-separated list of constructor types; it expects "allArgs" and "builder" (default "allArgs")
  -output string
        [optional] output file name (default "srcdir/<type>_gen.go")
  -type string
        [mandatory] a type name
  -version
        [optional] show the version information
  -withGetter
        [optional] generate a constructor along with getter functions for each field
```

## Motivation

Data encapsulation is a good practice to make software, and it is necessary to clearly indicate the boundary of the structure by controlling the accessibility of the data fields (i.e. private or public) for that. Basically keeping the data fields be private and immutable would be good to make software be robust because it can avoid unexpected field changing.

Golang has a simple way to do that by choosing the initial character's type: upper case or lower case. Once it has decided to use a field as private, it needs to make something like a constructor function, but golang doesn't have a mechanism to support constructor now.

Therefore this project aims to automatically generate constructors to use structures with private and immutable, easily.

## Pre requirements to run

- [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)

## Synopsis

### Generate all args constructor

1. write a struct type with `go:generate`

e.g.

```go
//go:generate gonstructor --type=Structure --constructorTypes=allArgs"
type Structure struct {
	foo string
	bar io.Reader
	Buz chan interface{}
}
```

2. execute `go generate ./...`
3. then `gonstructor` generates a constructor code

e.g.

```go
func NewStructure(foo string, bar io.Reader, buz chan interface{}) *Structure {
	return &Structure{foo: foo, bar: bar, Buz: buz}
}
```

### Generate builder (builder means GoF pattern's one)

1. write a struct type with `go:generate`

e.g.

```go
//go:generate gonstructor --type=Structure --constructorTypes=builder"
type Structure struct {
	foo string
	bar io.Reader
	Buz chan interface{}
}
```

2. execute `go generate ./...`
3. then `gonstructor` generates a buildr code

e.g.

```go
func NewStructureBuilder() *StructureBuilder {
	return &StructureBuilder{}
}
func (b *StructureBuilder) Foo(foo string) *StructureBuilder {
	b.foo = foo
	return b
}
func (b *StructureBuilder) Bar(bar io.Reader) *StructureBuilder {
	b.bar = bar
	return b
}
func (b *StructureBuilder) Buz(buz chan interface{}) *StructureBuilder {
	b.buz = buz
	return b
}
func (b *StructureBuilder) Build() *Structure {
	return &Structure{foo: b.foo, bar: b.bar, Buz: b.buz}
}
```

## How to ignore to contain a field in a constructor

`gonstructor:"-"` supports that.

e.g.

```go
type Structure struct {
	foo string
	bar int64 `gonstructor:"-"`
}
```

The generated code according to the above structure doesn't contain `bar` field.

## Where are pre-built binaries

[Releases](https://github.com/moznion/gonstructor/releases)

## How to build binaries

```
$ make all VERSION=x.y.z
```

## Author

moznion (<moznion@gmail.com>)

