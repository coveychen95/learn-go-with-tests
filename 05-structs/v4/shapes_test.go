package structs

import (
	"testing"

	structs "github.com/coveychen95/learn-go-with-tests/05-structs/v3"
)

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := structs.Rectangle{Width: 12.0, Height: 6.0}
		got := structs.Area(rectangle)
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := Area(circle)
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
