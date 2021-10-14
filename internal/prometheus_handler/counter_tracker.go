package prometheus_handler

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Counter stores the number of happenings of an event.
type CounterTracker struct {
	countVec *prometheus.CounterVec
	labels  prometheus.Labels
}

var (
	Counter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "number_of_events",
		Help: "Counting total number of some events like time outs",
	}, []string{"service_name", "function", "status"})
)

// NewResponseTimeTracker creates a new Tracker struct, which measures the response time of a function.
func NewCounterTracker(serviceName, fn, status string) *CounterTracker {
	return &CounterTracker{Counter,
		prometheus.Labels{"service_name": serviceName, "function": fn, "status": status}}
}

func (m *CounterTracker) Increase() {
	m.countVec.With(m.labels).Inc()
}

func (m *CounterTracker) SetStatus(status string) {
	m.labels["status"] = status
}