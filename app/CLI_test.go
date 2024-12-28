package poker_test

import (
	"strings"
	"testing"

	poker "github.com/biggsean/learn-go-with-tests2/app"
)

func TestCLI(t *testing.T) {
	t.Run("record sean win from user input", func(t *testing.T) {
		in := strings.NewReader("Sean wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		want := "Sean"

		poker.AssertPlayerWin(t, playerStore, want)
	})
	t.Run("record xavier win from user input", func(t *testing.T) {
		in := strings.NewReader("Xavier wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		want := "Xavier"

		poker.AssertPlayerWin(t, playerStore, want)
	})
}
