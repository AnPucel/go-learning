package maps

import "testing"

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	definition := "this is just a test"
	word := "test"

	dictionary.Add(word, definition)
	got, err := dictionary.Search(word)

	t.Run("new word", func(t *testing.T) {
		if err != nil {
			t.Fatal("should find added word:", err)
		}

		assertStrings(t, got, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		err = dictionary.Add(word, definition)

		assertError(t, err, ErrAlreadyExist)
	})
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
