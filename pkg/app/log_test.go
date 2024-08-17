package app

import (
	"testing"

	"github.com/whale1017/go-app/v10/pkg/errors"
	"github.com/whale1017/go-app/v10/pkg/logs"
)

func TestLog(t *testing.T) {
	DefaultLogger = t.Logf
	Log("hello", "world")
	Logf("hello %v", "Maxoo")
}

func TestServerLog(t *testing.T) {
	testSkipWasm(t)
	testLogger(t, serverLog)
}

func TestClientLog(t *testing.T) {
	testSkipNonWasm(t)
	testLogger(t, clientLog)
}

func testLogger(t *testing.T, l func(string, ...any)) {
	utests := []struct {
		scenario string
		value    any
	}{
		{
			scenario: "log",
			value:    logs.New("test").WithTag("type", "log"),
		},
		{
			scenario: "error",
			value:    errors.New("test").WithTag("type", "error"),
		},
	}

	for _, u := range utests {
		t.Run(u.scenario, func(t *testing.T) {
			l("%v", u.value)
		})
	}
}
