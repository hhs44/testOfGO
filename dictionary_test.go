package main

import "testing"

func TestSearch(t *testing.T) {

		assertStrings := func(t *testing.T, got , want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"
	assertStrings(t, got, want)
}

func Search(dictionary map[string]string, s string) string {
	return dictionary[s]
}
