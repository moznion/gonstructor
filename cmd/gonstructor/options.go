package main

import (
	"flag"
	"fmt"
	"strings"
)

type stringArrayFlag []string

func (a *stringArrayFlag) String() string        { return strings.Join(*a, ", ") }
func (a *stringArrayFlag) Set(item string) error { *a = append(*a, item); return nil }

type options struct {
	typeNames                                                      stringArrayFlag
	output, initFunc, getterPrefix, setterPrefix, batch            string
	constructorTypes                                               []string
	showVersion, withGetter, propagateInitFuncReturns, returnValue bool
	patterns, rawArgs                                              []string
}

func parseOptions(args []string) (options, error) {
	var o options
	flags := flag.NewFlagSet("gonstructor", flag.ContinueOnError)
	flags.Var(&o.typeNames, "type", "Type name; repeat to generate multiple types into one output.")
	flags.StringVar(&o.output, "output", "", "Output filename (default: <type>_gen.go).")
	constructorTypes := flags.String("constructorTypes", allArgsConstructorType, "Comma-separated constructor types: allArgs,builder.")
	flags.BoolVar(&o.showVersion, "version", false, "Show version information.")
	flags.BoolVar(&o.withGetter, "withGetter", false, "Generate getter methods.")
	flags.StringVar(&o.initFunc, "init", "", "Initialization method name.")
	flags.StringVar(&o.getterPrefix, "getterPrefix", "Get", "Getter method prefix.")
	flags.BoolVar(&o.propagateInitFuncReturns, "propagateInitFuncReturns", false, "Propagate initialization method return values.")
	flags.BoolVar(&o.returnValue, "returnValue", false, "Return a value instead of a pointer.")
	flags.StringVar(&o.setterPrefix, "setterPrefix", "", "Builder setter prefix.")
	flags.StringVar(&o.batch, "batch", "", "JSON batch manifest; paths are relative to the manifest.")
	if err := flags.Parse(args); err != nil {
		return o, err
	}
	o.rawArgs = append([]string(nil), args...)
	o.patterns = flags.Args()
	if len(o.patterns) == 0 {
		o.patterns = []string{"."}
	}
	if o.batch != "" {
		count := 0
		flags.Visit(func(*flag.Flag) { count++ })
		if count != 1 || flags.NArg() != 0 {
			return o, fmt.Errorf("-batch cannot be combined with other flags or package arguments")
		}
		return o, nil
	}
	if o.showVersion {
		return o, nil
	}
	if len(o.typeNames) == 0 {
		return o, fmt.Errorf("-type is required")
	}
	if len(o.typeNames) > 1 && o.output == "" {
		return o, fmt.Errorf("-output is required with multiple -type flags")
	}
	var err error
	o.constructorTypes, err = getConstructorTypes(*constructorTypes)
	return o, err
}
