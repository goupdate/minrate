# Rate Limiter

A simple and flexible rate limiter implementation in Go that allows you to control the rate at which actions are performed. This package is useful for scenarios where you need to limit the number of operations or requests over a specific duration, such as API calls, task executions, or event handling.

## Features

- **Configurable Rate Limits:** Set the number of actions and the duration for which the limit applies.
- **Flexible Duration:** Use any duration (e.g., seconds, minutes) for rate limiting.
- **Concurrency Safe:** Designed to work safely with goroutines.

## Installation

To use the `minrate` package in your Go project, you can install it using `go get`:

```sh
go get github.com/goupdate/minrate
```

## Usage

Here's a quick example of how to use the RateLimiter:

```
package main

import (
    "fmt"
    "github.com/goupdate/minrate"
    "time"
)

func main() {
    // Create a RateLimiter that allows 10 actions per minute
    rl := minrate.New(10, time.Minute)

    // Perform 15 actions
    for i := 0; i < 15; i++ {
        go func(i int) {
            rl.Wait() // Wait until it's allowed to perform the action
            b := rl.Can()
            fmt.Printf("Action %d, can: %t", i, b)
        }(i)
    }

    // Wait for all actions to complete
    time.Sleep(2 * time.Minute)
}
```

## API
New(actionsPerDuration int, duration time.Duration) *RateLimiter

Creates a new RateLimiter instance that allows actionsPerDuration actions over the specified duration.

    - actionsPerDuration: Number of actions allowed per duration.
    - duration: Time duration over which the actions are limited.

## Wait()

Blocks until an action can be performed according to the rate limit. Call this method before performing an action to ensure compliance with the rate limit.

## Can()

Informs will Wait() be locking right now or not. Call this method to test this case.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have suggestions or improvements.