package errorgrouptomb

import (
	"gopkg.in/tomb.v2"
	"testing"
)

func TestTomb(t *testing.T) {
	var to tomb.Tomb
	to.Dying()
}
