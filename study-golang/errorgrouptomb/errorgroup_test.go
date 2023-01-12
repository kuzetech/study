package errorgrouptomb

import (
	"golang.org/x/sync/errgroup"
	"testing"
)

func TestErrorGroup(t *testing.T) {
	var eg errgroup.Group
	eg.Wait()
}
