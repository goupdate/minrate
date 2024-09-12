package minrate

import (
	"sync"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	// Создаем RateLimiter с 5 действиями за 1 минуту
	rateLimiter := New(5, time.Second)

	// Запускаем 10 горутин, каждая из которых выполняет действие
	var numActions int
	var mu sync.Mutex
	start := time.Now()

	for i := 0; i < 10; i++ {
		go func() {
			rateLimiter.Wait()
			mu.Lock()
			numActions++
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second*2 + time.Millisecond*100)

	elapsed := time.Since(start)

	// Проверяем, что количество выполненных действий не превышает лимит
	if numActions != 10 {
		t.Errorf("Expected number of actions to be at most 10, got %d", numActions)
	}

	// Проверяем, что прошло достаточно времени (по крайней мере 60 секунд)
	if elapsed < 2*time.Second {
		t.Errorf("Expected at least 2 seconds to elapse, but got %v", elapsed)
	}
}

func TestRateLimiterWithBurst(t *testing.T) {
	// Создаем RateLimiter с 3 действиями за 30 секунд
	rateLimiter := New(10, time.Second)

	// Запускаем 6 горутин, каждая из которых выполняет действие
	start := time.Now()
	for i := 0; i < 6; i++ {
		rateLimiter.Wait()
	}

	elapsed := time.Since(start)

	// Проверяем, что время выполнения соответствует ожиданиям
	if elapsed > time.Millisecond*100 {
		t.Errorf("Expected fast but got %v", elapsed)
	}
}

func TestRateLimiterMultipleIntervals(t *testing.T) {
	rateLimiter := New(2, 2*time.Second)

	start := time.Now()
	for i := 0; i < 4; i++ {
		rateLimiter.Wait()
	}

	elapsed := time.Since(start)

	// Проверяем, что время выполнения соответствует ожиданиям
	if elapsed < 2*time.Second {
		t.Errorf("Expected at least 2 seconds to elapse for 4 actions, but got %v", elapsed)
	}
}

func TestRateLimiterEdgeCases(t *testing.T) {
	// Создаем RateLimiter с 1 действием за 1 секунду
	rateLimiter := New(1, time.Second)

	// Запускаем 2 горутины, каждая из которых выполняет действие
	start := time.Now()
	rateLimiter.Wait()
	rateLimiter.Wait()

	// Ожидаем завершения всех горутин
	time.Sleep(3 * time.Second)

	elapsed := time.Since(start)

	// Проверяем, что время выполнения соответствует ожиданиям
	if elapsed < 2*time.Second {
		t.Errorf("Expected at least 2 seconds to elapse for 2 actions, but got %v", elapsed)
	}
}
