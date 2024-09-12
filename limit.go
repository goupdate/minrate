package minrate

import (
	"time"
)

// RateLimiter представляет собой структуру для ограничения частоты действий
type RateLimiter struct {
	actionsPerDuration int
	duration           time.Duration
	tokens             chan struct{}
}

// New создает новый RateLimiter с заданным количеством действий и продолжительностью
func New(actionsPerDuration int, duration time.Duration) *RateLimiter {
	rl := &RateLimiter{
		actionsPerDuration: actionsPerDuration,
		duration:           duration,
		tokens:             make(chan struct{}, actionsPerDuration),
	}

	go rl.refillTokens()

	return rl
}

// refillTokens наполняет канал токенами с учетом установленного ограничения
func (rl *RateLimiter) refillTokens() {
	ticker := time.NewTicker(rl.duration)
	defer ticker.Stop()

	fill := func() {
		// Наполняем канал токенами
		for i := 0; i < rl.actionsPerDuration; i++ {
			select {
			case rl.tokens <- struct{}{}:
			default:
				// Канал уже заполнен
			}
		}
	}

	//init
	fill()

	for {
		select {
		case <-ticker.C:
			fill()
		}
	}
}

// Wait ожидает, пока действие станет возможным согласно ограничениям
func (rl *RateLimiter) Wait() {
	<-rl.tokens
}
