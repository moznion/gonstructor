package internal

import (
	"encoding/json"
	"fmt"
)

var ver string
var rev string

// ShowVersion shows the version and revision as JSON string.
func ShowVersion() {
	versionJSON, _ := json.Marshal(map[string]string{
		"version":  ver,
		"revision": rev,
	})
	fmt.Printf("%s\n", versionJSON)
}
