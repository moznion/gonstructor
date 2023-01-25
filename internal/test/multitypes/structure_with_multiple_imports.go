package multitypes

import (
	"io"
	"net"
)

// NOTE:
// This structure file is for checking whether the compilation passes for the generated file according to this.
// ref: https://github.com/moznion/gonstructor/pull/21
//
// Until the above pull request, the generated code hadn't passed the compilation.

//go:generate sh -c "$(cd ./\"$(git rev-parse --show-cdup)\" || exit; pwd)/dist/gonstructor_test --type=StructureWhichNeedsImport1 --type=StructureWhichNeedsImport2 --output structure_with_multiple_imports_gen.go --constructorTypes=allArgs,builder --withGetter"

type StructureWhichNeedsImport1 struct {
	foo net.Conn
}

type StructureWhichNeedsImport2 struct {
	bar io.Reader
}
