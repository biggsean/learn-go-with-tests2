package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	var person = Person{
		"Sean",
		Profile{48, "Sterling"},
	}

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{person.Name},
			[]string{"Sean"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{person.Name, person.Profile.City},
			[]string{"Sean", "Sterling"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{person.Name, person.Profile.Age},
			[]string{"Sean"},
		},
		{
			"nested fields",
			person,
			[]string{"Sean", "Sterling"},
		},
		{
			"pointers to things",
			&person,
			[]string{"Sean", "Sterling"},
		},
		{
			"slices",
			[]Profile{
				person.Profile,
				{33, "Herndon"},
			},
			[]string{"Sterling", "Herndon"},
		},
		{
			"arrays",
			[2]Profile{
				person.Profile,
				{33, "Herndon"},
			},
			[]string{"Sterling", "Herndon"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Bah",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Bah")
	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- person.Profile
			aChannel <- Profile{33, "Herndon"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Sterling", "Herndon"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return person.Profile, Profile{33, "Herndon"}
		}

		var got []string
		want := []string{"Sterling", "Herndon"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
