package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	OnlineFailureCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "online_failure_count",
		Help: "online failure message count",
	}, []string{"Index"})

	OnlineBuildDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "online_dealer_duration",
		Help: "online dealer build duration in milliseconds",
	}, []string{"Index"})

	OfflineFailureCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "offline_failure_count",
		Help: "offline failure message count",
	}, []string{"Index"})

	OfflineBuildDuration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "offline_dealer_duration",
			Help: "offline dealer build duration in milliseconds",
		}, []string{"Index"})
)

func init() {
	prometheus.MustRegister(OnlineFailureCount)
	prometheus.MustRegister(OnlineBuildDuration)
	prometheus.MustRegister(OfflineFailureCount)
	prometheus.MustRegister(OfflineBuildDuration)
}

func AvgTimerWithLabels(s *prometheus.SummaryVec, i int) func(...string) {
	start := time.Now()
	return func(l ...string) {
		total := time.Since(start)
		s.WithLabelValues(l...).Observe(float64(total) / float64(i) * 1e-6)
	}
}


func TimerWithLabels(s *prometheus.SummaryVec) func(...string) {
	start := time.Now()
	return func(l ...string) {
		total := time.Since(start)
		s.WithLabelValues(l...).Observe(float64(total))
	}
}

func main()  {

	moniter := TimerWithLabels(OnlineBuildDuration)
	for i := 1; i < 10 ; i ++ {
		time.Sleep(time.Second * time.Duration(i))
		OnlineFailureCount.WithLabelValues("test").Add(float64(i))
	}
	moniter("test")
}
