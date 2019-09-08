package matrix

import (
	"strings"
	"testing"
)

func TestMatrix_String(t *testing.T) {
	type entry struct {
		x, y int
		msg  string
	}

	tests := []struct {
		name    string
		want    string
		entries []entry
	}{
		{
			name: "output",
			entries: []entry{
				{
					x:   2,
					y:   2,
					msg: "hello",
				},
				{
					x:   4,
					y:   4,
					msg: "world",
				},
			},
			want: " \n" +
				" \n" +
				"  hello\n" +
				" \n" +
				"    world",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			matrix := Matrix{}
			for _, e := range test.entries {
				matrix.Put(e.x, e.y, e.msg)
			}

			output := matrix.String()
			if output != test.want {
				t.Errorf("Want output to be '%v' but got '%v'", strings.Replace(test.want, "\n", "\\n", -1), strings.Replace(output, "\n", "\\n", -1))
			}
		})
	}
}
