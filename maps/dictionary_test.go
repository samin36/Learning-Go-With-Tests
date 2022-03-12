package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("valid, existing key", func(t *testing.T) {
		dict := StringDict{"k1": "v1"}

		got, err := dict.Search("k1")
		want := "v1"

		assertNoError(t, err)
		assertStringEquals(t, got, want)
	})

	t.Run("invalid, nonexisting key", func(t *testing.T) {
		dict := StringDict{}

		got, err := dict.Search("k1")

		assertError(t, err, ErrKeyNotFound)
		assertStringEquals(t, got, "")
	})
}

func TestAdd(t *testing.T) {
	t.Run("add a new key-value pair", func(t *testing.T) {
		dict := StringDict{}

		addErr := dict.Add("k1", "v1")

		want := "v1"
		got, searchErr := dict.Search("k1")

		assertNoError(t, addErr)
		assertNoError(t, searchErr)
		assertStringEquals(t, got, want)
	})

	t.Run("add a key-value pair that already exists", func(t *testing.T) {
		dict := StringDict{"k1": "v1"}

		addErr := dict.Add("k1", "new_v1")

		want := "v1"
		got, searchErr := dict.Search("k1")

		assertError(t, addErr, ErrKeyExists)
		assertNoError(t, searchErr)
		assertStringEquals(t, got, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update an existing key-value pair", func(t *testing.T) {
		dict := StringDict{"k1": "v1"}

		updateErr := dict.Update("k1", "new_v1")

		want := "new_v1"
		got, searchErr := dict.Search("k1")

		assertNoError(t, updateErr)
		assertNoError(t, searchErr)
		assertStringEquals(t, got, want)
	})

	t.Run("update a non-existing key-value pair", func(t *testing.T) {
		dict := StringDict{}

		updateErr := dict.Update("k1", "v1")

		want := ""
		got, searchErr := dict.Search("k1")

		assertError(t, updateErr, ErrKeyNotFound)
		assertError(t, searchErr, ErrKeyNotFound)
		assertStringEquals(t, got, want)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete an existing key-value pair", func(t *testing.T) {
		dict := StringDict{"k1": "v1"}

		deleteErr := dict.Delete("k1")

		want := ""
		got, searchErr := dict.Search("k1")

		assertNoError(t, deleteErr)
		assertError(t, searchErr, ErrKeyNotFound)
		assertStringEquals(t, got, want)
	})

	t.Run("delete a non-existing key-value pair", func(t *testing.T) {
		dict := StringDict{}

		deleteErr := dict.Delete("k1")

		want := ""
		got, searchErr := dict.Search("k1")

		assertError(t, deleteErr, ErrKeyNotFound)
		assertError(t, searchErr, ErrKeyNotFound)
		assertStringEquals(t, got, want)
	})
}

func assertError(t testing.TB, gotErr, expectedErr error) {
	t.Helper()

	if gotErr != expectedErr {
		t.Errorf("expected error %q but received %q", gotErr, expectedErr)
	}
}

func assertNoError(t testing.TB, err error) {
	assertError(t, err, nil)
}

func assertStringEquals(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q and wanted %q", got, want)
	}
}
