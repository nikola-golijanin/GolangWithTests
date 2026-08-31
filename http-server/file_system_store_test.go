package main

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("league from a reader sorted", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer removeFile()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer removeFile()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
	})

	//file_system_store_test.go
	t.Run("store wins for existing players", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer removeFile()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer removeFile()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, removeFile := createTempFile(t, "")
		defer removeFile()

		_, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)
	})

	//file_system_store_test.go
	t.Run("league sorted", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer removeFile()

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}
	return tmpFile, removeFile
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
