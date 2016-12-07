package app_context

import (
	"log"
	"testing"
)

func TestStatsSender(t *testing.T) {
	app_ctx, err := NewAppContext("stats_test")
	if err != nil {
		log.Fatal(err)
	}

	err = app_ctx.StartStatsSender()
	if err != nil {
		t.Errorf("Starting stats sender failed: %+v", err)
	}

	err = app_ctx.StartStatsSender()
	if err == nil {
		t.Error("Starting stats sender succeeded after already starting")
	}

	err = app_ctx.StopStatsSender()
	if err != nil {
		t.Errorf("Stopping stats sender failed when running: %+v", err)
	}

	err = app_ctx.StopStatsSender()
	if err == nil {
		t.Error("Stopping stats a 2nd time succeeded when it shouldn't")
	}

	err = app_ctx.StartStatsSender()
	if err != nil {
		t.Errorf("Starting stats sender failed: %+v", err)
	}

	err = app_ctx.StopStatsSender()
	if err != nil {
		t.Errorf("Stopping stats sender failed when running: %+v", err)
	}
}
