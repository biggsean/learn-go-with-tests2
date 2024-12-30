package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"

	poker "github.com/biggsean/learn-go-with-tests2/app"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}

// var dummyStdIn = &bytes.Buffer{}
var dummyStdout = &bytes.Buffer{}

type GameSpy struct {
	StartCalled    bool
	StartedWith    int
	BlindAlert     []byte
	FinishedCalled bool
	FinishedWith   string
}

func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
	out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
	g.FinishedCalled = true
}

func TestCLI(t *testing.T) {
	t.Run("record sean win from user input", func(t *testing.T) {
		stdout := &bytes.Buffer{}

		in := userSends("2", "Sean")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 2)
		assertGameFinishCalledWith(t, game, "Sean")
	})
	t.Run("record xavier win from user input", func(t *testing.T) {
		in := userSends("5", "Xavier")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdout, game)
		cli.PlayPoker()

		assertGameStartedWith(t, game, 5)
		assertGameFinishCalledWith(t, game, "Xavier")
	})
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("7")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 7)
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("foo")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should have not started")
		}

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func assertGameStartedWith(t testing.TB, game *GameSpy, numberOfPlayers int) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartedWith == numberOfPlayers
	})
	if !passed {
		t.Errorf("expected game to be started with %d players but got %d", numberOfPlayers, game.StartedWith)
	}
}

func assertGameFinishCalledWith(t testing.TB, game *GameSpy, name string) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishedWith == name
	})

	if !passed {
		t.Errorf("expected game to be finished with %v players but got %v", name, game.FinishedWith)
	}
}

func assertMessageSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func userSends(userInputs ...string) *strings.Reader {
	return strings.NewReader(strings.Join(userInputs, "\n"))
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}
