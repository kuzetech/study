package sortslice

import "testing"

func Test_sortList(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortList()
		})
	}
}
