package constructor

import (
	"testing"
)

func TestToLowerCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "", want: ""},
		{input: "ID", want: "id"},
		{input: "utf8", want: "utf8"},
		{input: "Utf8", want: "utf8"},
		{input: "UTF8", want: "utf8"},
		{input: "utf8name", want: "utf8name"},
		{input: "utf8_name", want: "utf8name"},
		{input: "Name", want: "name"},
		{input: "name", want: "name"},
		{input: "NAME", want: "name"},
		{input: "UserName", want: "userName"},
		{input: "userName", want: "userName"},
		{input: "user_Name", want: "userName"},
		{input: "MoZnIoN", want: "moZnIoN"},
	}

	for _, test := range tests {
		got := toLowerCamel(test.input)
		if got != test.want {
			t.Errorf("toLowerCamel: want %v, but %v for %v:", test.want, got, test.input)
		}
	}
}
