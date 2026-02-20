package taxonomy

import (
	"fmt"
	"testing"
)

func TestTypeTaxonomy_String(t *testing.T) {
	tests := []struct {
		name   string
		input  TypeTaxonomy
		output string
	}{
		{
			name:   "test animals",
			input:  Animals,
			output: "Animals",
		},
		{
			name:   "test animals",
			input:  Plants,
			output: "Plants",
		},
		{
			name:   "test animals",
			input:  None,
			output: "None",
		}, {
			name:   "test animals",
			input:  Tools,
			output: "Tools",
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fmt.Sprint(tt.input)
			if got != tt.output {
				t.Errorf("TestTypeTaxonomy_String() got = %v, want %v", got, tt.output)
			}
		})
	}
}
