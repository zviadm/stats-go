package datadog

import (
	"context"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zviadm/stats-go/metrics"
)

// To enable this test, create `datadog.key` file and place your API key in it. Note that
// this test will log real stats to your DataDog account.
func TestExporter(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	apiKeyB, err := ioutil.ReadFile("datadog.key")
	if err != nil {
		t.SkipNow()
	}
	apiKey := string(apiKeyB)
	apiKey = strings.TrimSpace(apiKey)

	cmd := exec.Command("/opt/datadog-dogstatsd/bin/dogstatsd", "start")
	cmd.Env = append(os.Environ(),
		"DD_API_KEY="+apiKey,
		"DD_ENABLE_METADATA_COLLECTION=false",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	require.NoError(t, err)
	defer cmd.Process.Kill()

	metrics.SetInstanceNameAndNodeTags("exporter_test", nil)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ExporterGo(ctx)

	c1 := metrics.DefineCounter("exporter_test/c1", metrics.WithTags("tag_one"))
	g1 := metrics.DefineGauge("exporters_test/g1")

	for idx := 0; idx < 3; idx++ {
		c1.V(metrics.KV{"tag_one": "t1"}).Count(float64(idx))
		g1.V().Set(float64(idx))
		time.Sleep(*flagPushFrequency)
	}
	cmd.Process.Signal(syscall.SIGTERM)
	err = cmd.Wait()
	require.NoError(t, err)
}
