package constructor

import (
	"fmt"
	"strings"

	g "github.com/moznion/gowrtr/generator"
)

func generateStructure(typeName string, keyValues []string, indentLevel int, returnValue bool) string {
	indent := g.BuildIndent(indentLevel)
	nextIndent := g.BuildIndent(indentLevel + 1)

	structure := fmt.Sprintf(
		"%s{\n%s%s,\n%s}",
		typeName,
		nextIndent,
		strings.Join(keyValues, ",\n"+nextIndent),
		indent,
	)
	return withConditionalPrefix(structure, "&", !returnValue)
}
