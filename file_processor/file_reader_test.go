package FileProcessor

import (
	"reflect"
	"testing"
)

func assertMatch(t *testing.T, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %s, Want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Wanted %s, got %s", want, got)
	}

}

func TestReadWebsiteList(t *testing.T) {
	t.Run("reads contents of file and converts to string list", func(t *testing.T) {
		got, _ := ReadFile("./dummy_data.txt")
		want := []string{"bookface.com", "twotter.com", "instasham.com"}

		assertMatch(t, got, want)
	})

	t.Run("works with other file", func(t *testing.T) {
		got, _ := ReadFile("./dummy_data2.txt")
		want := []string{"foo.com", "bar.ca", "baz.co.uk"}

		assertMatch(t, got, want)
	})

	t.Run("errors when file is not found", func(t *testing.T) {
		_, err := ReadFile("./some_missing_file.txt")
		want := ReaderErr("File ./some_missing_file.txt not found")

		assertError(t, err, want)
	})
}
