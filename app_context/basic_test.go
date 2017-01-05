package app_context

import (
	"log"
	"os"
	"testing"
)

func TestTiltEnv(t *testing.T) {
	os.Unsetenv("TILT_ENVIRONMENT")
	app_ctx, err := NewAppContext("basic_test")
	if err != nil {
		log.Fatal(err)
	}
	if tiltEnv := app_ctx.TiltEnv(); tiltEnv != "development" {
		t.Errorf("environment is not development: %s", tiltEnv)
	}

	os.Setenv("TILT_ENVIRONMENT", "testing")

	app_ctx, err = NewAppContext("basic_test")
	if err != nil {
		log.Fatal(err)
	}
	if tiltEnv := app_ctx.TiltEnv(); tiltEnv != "testing" {
		t.Errorf("environment is not testing: %s", tiltEnv)
	}

	os.Setenv("TILT_ENVIRONMENT", "derp")

	app_ctx, err = NewAppContext("basic_test")
	if err == nil {
		t.Error("app context should have failed")
	}
}

func TestCodeVersion(t *testing.T) {
	os.Unsetenv("CODE_VERSION")
	app_ctx, err := NewAppContext("basic_test")
	if err != nil {
		log.Fatal(err)
	}
	if vers := app_ctx.CodeVersion(); vers != "" {
		t.Errorf("Code version is not empty: %s", vers)
	}

	os.Setenv("CODE_VERSION", "abczyx")

	app_ctx, err = NewAppContext("basic_test")
	if err != nil {
		log.Fatal(err)
	}
	if vers := app_ctx.CodeVersion(); vers != "abczyx" {
		t.Errorf("Code version is not 'abczyx' %s", vers)
	}
}
