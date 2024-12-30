package poker_test

import (
	"fmt"
	"io"
	"testing"
	"time"

	poker "github.com/biggsean/learn-go-with-tests2/app"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on a game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)
		game.Start(5, io.Discard)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
	t.Run("schedules alerts on a game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)
		game.Start(7, io.Discard)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
			{48 * time.Minute, 500},
			{60 * time.Minute, 600},
			{72 * time.Minute, 800},
			{84 * time.Minute, 1000},
			{96 * time.Minute, 2000},
			{108 * time.Minute, 4000},
			{120 * time.Minute, 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	t.Run("record the winner of the game", func(t *testing.T) {
		var tests = []struct {
			name string
		}{
			{"Sean"},
			{"Shane"},
			{"Xavier"},
			{"Owen"},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				store := &poker.StubPlayerStore{}
				game := poker.NewTexasHoldem(dummyBlindAlerter, store)

				game.Finish(tt.name)
				poker.AssertPlayerWin(t, store, tt.name)
			})
		}
	})
}

func checkSchedulingCases(cases []poker.ScheduledAlert, t *testing.T, blindAlerter *poker.SpyBlindAlerter) {
	t.Helper()
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()

	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}
	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}
}
