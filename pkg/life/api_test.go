package life_test

import (
	"testing"

	"github.com/jaztec/golife/pkg/life"
)

func TestNewGrid(t *testing.T) {
	var expect int32 = 196
	t.Run("test exact number of cells", func(t *testing.T) {
		_, got := life.NewGrid(expect)

		if got != expect {
			t.Errorf("expected %d cells but %d where created", expect, got)
		}
	})
	t.Run("test calculated number of cells", func(t *testing.T) {
		_, got := life.NewGrid(200)

		if got != expect {
			t.Errorf("expected %d cells but %d where created", expect, got)
		}
	})
}
