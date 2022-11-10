package traverseDir

import "testing"

func Test_traverseDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test", args: args{path: "./"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			traverseDir(tt.args.path)
		})
	}
}
