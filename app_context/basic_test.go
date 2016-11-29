package app_context

import (
	"log"
	"os"
	"testing"
)

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
