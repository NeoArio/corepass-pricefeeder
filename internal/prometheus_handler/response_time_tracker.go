package prometheus_handler

import (
"github.com/prometheus/client_golang/prometheus"
"github.com/prometheus/client_golang/prometheus/promauto"
"time"
)

// Tracker store data we need to create a histogram.
type Tracker struct {
	histVec *prometheus.HistogramVec
	start   time.Time
	labels  prometheus.Labels
	Duration float64
}

var (
	ResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "response_time_seconds",
		Help: "Response time of methods",
		//Buckets: prometheus.LinearBuckets(0.5, 0.5, 20),
		Buckets: []float64{0.05, 0.1, 0.2, 0.3, 0.4, 0.5, 0.650, 0.8, 1, 1.2, 1.5, 1.8, 2.1, 2.5, 3, 4},
	}, []string{"service_name", "function", "status"})

)

// NewResponseTimeTracker creates a new Tracker struct, which measures the response time of a function.
func NewResponseTimeTracker(serviceName, fn, status string) *Tracker {
	return &Tracker{ResponseTime, time.Now(),
		prometheus.Labels{"service_name": serviceName, "function": fn, "status": status},0}
}


func (m *Tracker) Observe() {
	m.histVec.With(m.labels).Observe(time.Since(m.start).Seconds())
	m.Duration = time.Since(m.start).Seconds()
}

func (m *Tracker) SetStatus(status string) {
	m.labels["status"] = status
}
