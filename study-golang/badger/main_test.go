package badger

import "testing"

func Test_openDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//openDB()
			//printAllKeys()
			//keyNotExist()
			keyExist()
		})
	}
}
