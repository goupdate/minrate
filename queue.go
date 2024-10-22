package minrate

import (
	"sync"
	"time"
)

func init() {
	go refillTokens()
}

type tqueue struct {
	sync.Mutex
	limiters []*RateLimiter
}

var queue tqueue

// refillTokens наполняет канал токенами с учетом установленного ограничения
func refillTokens() {
	ticker := time.NewTicker(time.Millisecond * 50)

	fill := func(t chan struct{}, actions int) {
		// Наполняем канал токенами
		for i := 0; i < actions; i++ {
			select {
			case t <- struct{}{}:
			default:
				// Канал уже заполнен
			}
		}
	}

	for tn := range ticker.C {
		queue.Lock()
		for _, q := range queue.limiters {
			if tn.Sub(q.lastFill) > q.duration {
				fill(q.tokens, q.actionsPerDuration)
				q.lastFill = tn
			}
		}
		queue.Unlock()
	}
}
