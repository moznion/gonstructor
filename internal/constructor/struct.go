package constructor

import (
	"fmt"
	"strings"

	g "github.com/moznion/gowrtr/generator"
)

func generateStructure(typeName string, keyValues []string, indentLevel int) string {
	indent := g.BuildIndent(indentLevel)
	nextIndent := g.BuildIndent(indentLevel + 1)

	return fmt.Sprintf(
		"&%s{\n%s%s,\n%s}",
		typeName,
		nextIndent,
		strings.Join(keyValues, ",\n"+nextIndent),
		indent,
	)
}
