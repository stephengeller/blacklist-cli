package internal

import (
	"errors"
	"testing"
)

type MockLogger struct {
	Calls int
}

func (m *MockLogger) Fatal(err error) {
	m.Calls++
}

func assertMatch(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d, Want %d", got, want)
	}
}

func TestCheck(t *testing.T) {
	t.Run("does a log.fatal if err is not nil", func(t *testing.T) {
		logger := MockLogger{}
		Check(&logger, errors.New("some err"))
		assertMatch(t, logger.Calls, 1)
	})

	t.Run("doesn't do a log.fatal if err is nil", func(t *testing.T) {
		logger := MockLogger{}
		Check(&logger, nil)
		assertMatch(t, logger.Calls, 0)
	})
}
