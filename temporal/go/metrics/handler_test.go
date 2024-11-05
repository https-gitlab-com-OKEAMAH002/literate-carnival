package metrics

import (
	"fmt"
	"math"
	"sync"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.skia.org/infra/go/metrics2"
)

func newClient() metrics2.Client {
	// This wipes out all previous metrics.
	prometheus.DefaultRegisterer = prometheus.NewRegistry()
	return metrics2.GetDefaultClient()
}

func TestHandler_NewHandler_AppendTags(t *testing.T) {
	tag1 := map[string]string{
		"some-key1": "some-value1",
	}
	tag2 := map[string]string{
		"some-key2": "some-value2",
	}
	all := map[string]string{
		"some-key1": "some-value1",
		"some-key2": "some-value2",
	}

	c := newClient()
	h1 := NewMetricsHandler(tag1, c)
	h2 := h1.WithTags(tag2)
	require.EqualValues(t, tag1, h1.tags)
	require.EqualValues(t, all, h2.(*metricsHandler).tags)
}

func TestHandler_WithSameCounter_ReturnSameCounter(t *testing.T) {
	const cn = "some_counter"
	h := NewMetricsHandler(map[string]string{}, newClient())
	require.Equal(t, h.Counter(cn), h.Counter(cn))
	require.NotEqual(t, h.Counter("other"), h.Counter(cn))
}

func TestHandler_WithSameGauge_ReturnSameGauge(t *testing.T) {
	const gn = "some_gauge"
	h := NewMetricsHandler(map[string]string{}, newClient())
	require.Equal(t, h.Gauge(gn), h.Gauge(gn))
	require.NotEqual(t, h.Gauge("other"), h.Gauge(gn))
}

func TestHandler_OnSameTimer_ReturnSameTimer(t *testing.T) {
	const tn = "some_timer"
	h := NewMetricsHandler(map[string]string{}, newClient())
	require.Equal(t, h.Timer(tn), h.Timer(tn))
	require.NotEqual(t, h.Timer("other"), h.Timer(tn))
}

func TestHandler_WithCounter_RecordMetric(t *testing.T) {
	const cn = "some_counter"
	h := NewMetricsHandler(map[string]string{}, newClient())
	h.Counter(cn).Inc(1)
	require.EqualValues(t, 1, h.Counter(cn).(metrics2.Counter).Get())
	require.NotEqualValues(t, 1, h.Counter("other").(metrics2.Counter).Get())
}

func TestHandler_WithGauge_RecordMetric(t *testing.T) {
	const gn = "some_gauge"
	h := NewMetricsHandler(map[string]string{}, newClient())
	h.Gauge(gn).Update(2.2)
	require.InDelta(t, 2.2, h.Gauge(gn).(metrics2.Float64Metric).Get(), 1e-9)
	require.Greater(t, math.Abs(2.2-h.Gauge("other").(metrics2.Float64Metric).Get()), 1e-9)
}

func TestHanlder_RecordMetrics_InParallel(t *testing.T) {
	var wg sync.WaitGroup
	const count = 500

	// Spawn many goroutines for 4 different transactions
	wg.Add(count * 4)

	h := NewMetricsHandler(map[string]string{}, newClient())
	// Read & Updates the metrics forward and backward.
	for i := 0; i < count; i++ {
		counter_forward := fmt.Sprintf("counter_%d", i)
		counter_backward := fmt.Sprintf("counter_%d", count-i-1)
		gauge_forward := fmt.Sprintf("gauge_%d", i)
		gauge_backward := fmt.Sprintf("gauge_%d", count-i-1)
		go func() {
			h.Counter(counter_forward).Inc(1)
			wg.Done()
		}()
		go func() {
			h.Counter(counter_backward).Inc(1)
			wg.Done()
		}()
		go func() {
			h.Gauge(gauge_forward).Update(2.0)
			wg.Done()
		}()
		go func() {
			h.Gauge(gauge_backward).Update(2.0)
			wg.Done()
		}()
	}
	wg.Wait()

	for i := 0; i < count; i++ {
		counter := fmt.Sprintf("counter_%d", i)
		gauge := fmt.Sprintf("gauge_%d", i)
		// Each counter gets increased twice.
		assert.EqualValues(t, 2, h.Counter(counter).(metrics2.Counter).Get(), "counter %d", i)
		assert.InDelta(t, 2.0, h.Gauge(gauge).(metrics2.Float64Metric).Get(), 1e-9, "gauge %d", i)
	}
}
