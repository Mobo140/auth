package breaker

import (
	"log"
	"time"

	"github.com/sony/gobreaker"
)

const (
	timeout         = 5 * time.Second
	maxRequests     = 3
	maxFailureRatio = 0.5
)

func Init() *gobreaker.CircuitBreaker {
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "auth-service",
		MaxRequests: maxRequests,
		Timeout:     timeout,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio >= maxFailureRatio
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("CircuitBreaker %s state change from %s to %s", name, from.String(), to.String())
		},
	})

	return cb
}
