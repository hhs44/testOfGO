package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Huang", "")
	want := "Hello, Huang"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello2(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// 声明方法是辅助函数
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
		//if got != want {
		//	t.Errorf("got %q want %q", got, want)
		//}
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

		//if got != want {
		//	t.Errorf("got %q want %q", got, want)
		//}
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("World", "French")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})
}
