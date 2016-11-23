package app_context

import (
	"log"
	"os"
	"testing"
)

func TestRollbarEnabled(t *testing.T) {
	os.Setenv("ROLLBAR_API_KEY", "FOO")
	os.Unsetenv("ROLLBAR_DISABLE")

	app_ctx, err := NewAppContext("rollbar_test")
	if err != nil {
		log.Fatal(err)
	}

	if !app_ctx.RollbarEnabled() {
		t.Error("No ROLLBAR_DISABLE env but rollbar is disabled")
	}

	os.Setenv("ROLLBAR_DISABLE", "false")

	app_ctx, err = NewAppContext("rollbar_test")
	if err != nil {
		log.Fatal(err)
	}

	if !app_ctx.RollbarEnabled() {
		t.Error("ROLLBAR_DISABLE=false env but rollbar is disabled")
	}

	os.Setenv("ROLLBAR_DISABLE", "true")

	app_ctx, err = NewAppContext("rollbar_test")
	if err != nil {
		log.Fatal(err)
	}

	if app_ctx.RollbarEnabled() {
		t.Error("ROLLBAR_DISABLE=true env but rollbar is not disabled")
	}

	os.Unsetenv("ROLLBAR_API_KEY")
	os.Unsetenv("ROLLBAR_DISABLE")

	app_ctx, err = NewAppContext("rollbar_test")
	if err != nil {
		log.Fatal(err)
	}

	if app_ctx.RollbarEnabled() {
		t.Error("No ROLLBAR_API_KEY but rollbar is enabled")
	}
}
