package FileProcessor

import (
	"log"
	"os"
	"testing"
)

const testFilePath = "./example_hostfile"

func ResetTestHostile(filename string) {
	err := os.Remove(filename)
	_, err = os.Create(filename)

	if err != nil {
		log.Fatal(err)
		return
	}

}

func AddTestEntriesForDeletion(t *testing.T, websites []string) error {
	err := AddLinesToFile("./example_hostfile", websites)
	assertError(t, err, nil)
	return err
}

func TestAddSitesToHostfile(t *testing.T) {
	t.Run("success scenario", func(t *testing.T) {
		ResetTestHostile(testFilePath)
		websites := []string{"foo.com", "bar.com", "baz.com"}
		err := AddLinesToFile(testFilePath, websites)

		assertError(t, err, nil)

		got, err := ReadFile("./example_hostfile")
		want := []string{"foo.com", "bar.com", "baz.com"}

		assertError(t, err, nil)
		assertMatch(t, got, want)
	})

	t.Run("missing file", func(t *testing.T) {
		ResetTestHostile(testFilePath)

		websites := []string{"foo.com", "bar.com", "baz.com"}
		err := AddLinesToFile("./missing_hostfile", websites)

		assertError(t, err, ReaderErr("File ./missing_hostfile not found"))
	})
}

func TestRemoveLinesFromFile(t *testing.T) {
	t.Run("success scenario", func(t *testing.T) {
		ResetTestHostile(testFilePath)
		websites := []string{"one_before", "foo.com", "bar.com", "baz.com", "one_after"}
		err := AddTestEntriesForDeletion(t, websites)
		got, err := ReadFile(testFilePath)
		assertError(t, err, nil)
		assertMatch(t, got, websites)

		sitesToRemove := []string{"foo.com", "bar.com", "baz.com"}

		err = RemoveLinesFromFile(testFilePath, sitesToRemove)
		assertError(t, err, nil)

		got, err = ReadFile(testFilePath)
		want := []string{"one_before", "one_after"}
		assertMatch(t, got, want)
		assertError(t, err, nil)
	})

	t.Run("other contents", func(t *testing.T) {
		ResetTestHostile(testFilePath)
		websites := []string{"blah", "blim", "bop", "baz.com", "one_after"}
		err := AddTestEntriesForDeletion(t, websites)
		got, err := ReadFile(testFilePath)
		assertError(t, err, nil)
		assertMatch(t, got, websites)

		sitesToRemove := []string{"bop", "blim", "baz.com"}

		err = RemoveLinesFromFile(testFilePath, sitesToRemove)
		assertError(t, err, nil)

		got, err = ReadFile(testFilePath)
		want := []string{"blah", "one_after"}
		assertMatch(t, got, want)
		assertError(t, err, nil)
	})
}

func TestReplaceFileContents(t *testing.T) {
	t.Run("success scenario", func(t *testing.T) {
		ResetTestHostile(testFilePath)
		websites := []string{"one_before", "foo.com", "bar.com", "baz.com", "one_after"}
		err := AddTestEntriesForDeletion(t, websites)
		got, err := ReadFile(testFilePath)
		assertError(t, err, nil)
		assertMatch(t, got, websites)

		err = ReplaceFileContents(testFilePath, []string{"foo", "bar"})
		assertError(t, err, nil)

		got, err = ReadFile(testFilePath)
		want := []string{"foo", "bar"}
		assertMatch(t, got, want)
		assertError(t, err, nil)
	})
}
