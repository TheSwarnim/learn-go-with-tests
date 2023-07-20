package racer

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(a, b string) string {
// 	aDuration := measureResponseTime(a)

// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

const tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{}) // empty struct occupies no memory
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
// We have defined a function ping which creates a chan struct{} and returns it.
// In our case, we don't care what type is sent to the channel, we just want to signal we are done and closing the channel works perfectly!
// Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool. Since we are closing and not sending anything on the chan, why allocate anything?
// Inside the same function, we start a goroutine which will send a signal into that channel once we have completed http.Get(url).

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }