package app_context

import (
	"log"
	"os"
	"testing"
)

func TestMetricsTagsEnv(t *testing.T) {
	os.Setenv("METRICS_TAGS", "k1=v,k2=v2,k3=a=b,k4")

	app_ctx, err := NewAppContext("metrics_test")
	if err != nil {
		log.Fatal(err)
	}

	tags := app_ctx.MetricsClient().GetTags()

	if len(tags) != 6 {
		t.Errorf("tags doesn't have 6 keys, it has %d: %+v", len(tags), tags)
	}

	if _, ok := tags["host"]; !ok {
		t.Errorf("tags[host] set not set: %+v", tags)
	}

	if tags["application"] != "metrics_test" {
		t.Errorf("tags[application] != metrics_test: %+v", tags)
	}

	if tags["k1"] != "v" {
		t.Errorf("tags[k1] != v: %+v", tags)
	}

	if tags["k2"] != "v2" {
		t.Errorf("tags[k2] != v2: %+v", tags)
	}

	if tags["k3"] != "a=b" {
		t.Errorf("tags[k3] != a=b: %+v", tags)
	}

	if tags["k4"] != "" {
		t.Errorf("tags[k4] != '': %+v", tags)
	}
}

func TestMetricsNamespace(t *testing.T) {
	os.Unsetenv("METRICS_NAMESPACE")

	app_ctx, err := NewAppContext("metrics_test")
	if err != nil {
		log.Fatal(err)
	}

	if s := app_ctx.MetricsClient().GetNamespace(); s != "metrics_test." {
		t.Errorf("Namespace is not metrics_test. with no Env: %s", s)
	}

	os.Setenv("METRICS_NAMESPACE", "custom_ns.")

	app_ctx, err = NewAppContext("metrics_test")
	if err != nil {
		log.Fatal(err)
	}

	if s := app_ctx.MetricsClient().GetNamespace(); s != "custom_ns." {
		t.Errorf("Namespace is not 'custom_ns.' from Env: %s", s)
	}
}

func TestMetricsEnabled(t *testing.T) {
	os.Unsetenv("METRICS_DISABLE")

	app_ctx, err := NewAppContext("metrics_test")
	if err != nil {
		log.Fatal(err)
	}

	if !app_ctx.MetricsEnabled() {
		t.Error("No METRICS_DISABLE env but metrics is disabled")
	}

	os.Setenv("METRICS_DISABLE", "false")

	app_ctx, err = NewAppContext("metrics_test")
	if err != nil {
		log.Fatal(err)
	}

	if !app_ctx.MetricsEnabled() {
		t.Error("METRICS_DISABLE=false env but metrics is disabled")
	}

	os.Setenv("METRICS_DISABLE", "true")

	app_ctx, err = NewAppContext("metrics_test")
	if err != nil {
		log.Fatal(err)
	}

	if app_ctx.MetricsEnabled() {
		t.Error("METRICS_DISABLE=true env but metrics is not disabled")
	}
}
