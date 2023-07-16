# gonstructor

A command-line tool to generate a constructor for the struct.

## Installation

```
$ go install github.com/moznion/gonstructor/cmd/gonstructor@latest
```

Also, you can get the pre-built binaries on [Releases](https://github.com/moznion/gonstructor/releases).

Or get it with [gobinaries.com](https://gobinaries.com):

```bash
curl -sf https://gobinaries.com/moznion/gonstructor | sh
```

## Dependencies

gonstructor depends on [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) for fixing import paths and formatting code, you need to install it:

```
$ go install golang.org/x/tools/cmd/goimports@latest
```

## Usage

```
Usage of gonstructor:
  -constructorTypes string
        [optional] comma-separated list of constructor types; it expects "allArgs" and "builder" (default "allArgs")
  -init string
        [optional] name of function to call on object after creating it
  -output string
        [optional] Output file name (default "srcdir/<type>_gen.go"). See also "-type" option's description.'
  -propagateInitFuncReturns
        [optional] If this option is specified, the generated constructor propagates the return values that come from the init function specified by the "-init" option, e.g. when the init function returns an "error" value, the generated constructor returns (*YourStructType, error). Known issue: If this option is used with the multiple --type options, probably it won't be the expected result.
  -type value
        [mandatory] A type name. It accepts this option occurs multiple times to output the generated code of the multi types into a single file. If this option is given multiple times, the "-output" option becomes mandatory.
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
//go:generate gonstructor --type=Structure --constructorTypes=allArgs
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
//go:generate gonstructor --type=Structure --constructorTypes=builder
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
type StructureBuilder struct {
	foo        string
	bar        io.Reader
	buz        chan interface{}
	bufferSize int
}

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

func (b *StructureBuilder) BufferSize(bufferSize int) *StructureBuilder {
	b.bufferSize = bufferSize
	return b
}

func (b *StructureBuilder) Build() *Structure {
	return &Structure{
		foo:        b.foo,
		bar:        b.bar,
		Buz:        b.buz,
		bufferSize: b.bufferSize,
	}
}
```

### Call a initializer

1. write a struct type with `go:generate`
2. write a function that initializes internal fields
3. pass its name to `-init` parameter

e.g.

```go
//go:generate gonstructor --type=Structure -init construct
type Structure struct {
	foo        string
	bar        io.Reader
	Buz        chan interface{}
	bufferSize int
	buffer     chan []byte `gonstructor:"-"`
}

func (structure *Structure) construct() {
	structure.buffer = make(chan []byte, structure.bufferSize)
}
```

2. execute `go generate ./...`
3. then `gonstructor` generates a buildr code

e.g.

```go
func NewStructure(
	foo string,
	bar io.Reader,
	buz chan interface{},
	bufferSize int,
) *Structure {
	r := &Structure{
		foo:        foo,
		bar:        bar,
		Buz:        buz,
		bufferSize: bufferSize,
	}

	r.construct()

	return r
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

## How to output the generated code of each type into a single file

This CLI tool can have the `--type` option multiple times, and it must have also `--output` option.

example:

```
//go:generate gonstructor --type=AlphaStructure --type=BravoStructure --constructorTypes=allArgs,builder --withGetter --output=./alpha_and_bravo_gen.go"
```

## How to propagate the returned values come from `-init` func

`-propagateInitFuncReturns` option supports that.

For example,

```go
//go:generate gonstructor --type=Struct --constructorTypes=allArgs,builder --init validate --propagateInitFuncReturns
type Struct struct {
	foo string
}

func (s *Struct) validate() error {
	// do something with the created `Struct` value.
	return err
}
```

then it generates the following Go code:

```go
func NewStruct(foo string) (*Struct, error) {
	r := &Struct{
		foo: foo,
	}

	ret_validate0 := r.validate()

	return r, ret_validate0
}
```

As you can see, the generated constructor `NewStruct()` returns the constructed value and an error that comes from `validate()` function.

## How to build binaries

Binaries are built and uploaded by [goreleaser](https://goreleaser.com/). Please refer to the configuration file: [.goreleaser.yml](./.goreleaser.yml)

## Author

moznion (<moznion@gmail.com>)
