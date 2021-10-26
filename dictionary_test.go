package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is a test")

		expected := "this is a test"

		compareErr(t, err, nil)
		compareDefinition(t, dictionary, word, expected)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definitition := "this is a test"
		dictionary := Dictionary{word: definitition}
		err := dictionary.Add(word, "this is a new test")

		compareErr(t, err, ErrorWordExists)
		compareDefinition(t, dictionary, word, definitition)
	})
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("search exists word", func(t *testing.T) {
		compareDefinition(t, dictionary, "test", "this is a test")
	})

	t.Run("search with nonexistent word", func(t *testing.T) {
		_, err := dictionary.Search("nonexistent")

		compareErr(t, err, ErrorNonExistentWord)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definitition := "this is a test"
	dictionary := Dictionary{word: definitition}
	newDefinition := "this is a test updated"

	t.Run("updating a existing word", func(t *testing.T) {
		err := dictionary.Update(word, newDefinition)

		compareErr(t, err, nil)
		compareDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("updating a nonexistent word", func(t *testing.T) {
		wordNonExistent := "teste2"
		err := dictionary.Update(wordNonExistent, newDefinition)

		compareErr(t, err, ErrorNonExistentWordToUpdate)
	})

}

func TestDestroy(t *testing.T) {
	word := "test"
	definitition := "this is a test"
	dictionary := Dictionary{word: definitition}

	dictionary.Destroy(word)

	_, err := dictionary.Search(word)

	if err != ErrorNonExistentWord {
		t.Errorf("expect '%s' to be deleted", word)
	}
}

func compareDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	result, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should have found added word:", err)
	}
	if result != definition {
		t.Errorf("result %s, expected %s", result, definition)
	}
}

func compareErr(t *testing.T, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("result error '%s', expected '%s'", result, expected)
	}

}
