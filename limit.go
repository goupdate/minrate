package minrate

import (
	"slices"
	"time"
)

// RateLimiter представляет собой структуру для ограничения частоты действий
type RateLimiter struct {
	actionsPerDuration int
	duration           time.Duration
	lastFill           time.Time
	tokens             chan struct{}
}

// New создает новый RateLimiter с заданным количеством действий и продолжительностью
func New(actionsPerDuration int, duration time.Duration) *RateLimiter {
	rl := &RateLimiter{
		actionsPerDuration: actionsPerDuration,
		duration:           duration,
		tokens:             make(chan struct{}, actionsPerDuration),
	}

	queue.Lock()
	queue.limiters = append(queue.limiters, rl)
	queue.Unlock()

	return rl
}

// Wait ожидает, пока действие станет возможным согласно ограничениям
func (rl *RateLimiter) Wait() {
	<-rl.tokens
}

// can now take action or wait?
// true if now wait i.e. Wait() call will end immidiatelly
func (rl *RateLimiter) Can() bool {
	return len(rl.tokens) > 0
}

// remove limiter from checking queue
func (rl *RateLimiter) Close() {
	queue.Lock()
	defer queue.Unlock()
	queue.limiters = slices.DeleteFunc(queue.limiters, func(r *RateLimiter) bool {
		return r == rl
	})
}
