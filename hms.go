package hms

import (
	"errors"
	"fmt"
	"time"
)

// HMS holds time in format HH:MM:SS with 24 hours limit. The value is stored in seconds.
type HMS int64

var ErrParseFailed = errors.New("invalid string format. Expected HH:MM:SS")

const day int64 = 24 * 60 * 60

// New builds HMS from time.Duration.
func New(d time.Duration) HMS {
	if d == 0 {
		return HMS(0)
	}
	h := int64(d.Seconds())
	if h >= day {
		h = h % day
	}
	return HMS(h)
}

// Parse receives string in format HH:MM:SS, parses it and returns HMS.
func Parse(str string) (HMS, error) {
	var h, m, s int64
	_, err := fmt.Sscanf(str, "%d:%d:%d", &h, &m, &s)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrParseFailed, err)
	}
	if m < 0 || m > 59 || s < 0 || s > 59 {
		return 0, fmt.Errorf("%w: invalid minutes or seconds", ErrParseFailed)
	}

	res := h*60*60 + m*60 + s
	if res >= day {
		return HMS(res % day), nil
	}

	return HMS(res), nil
}

// Add adds duration to HMS and returns new HMS.
// If the result of addition causes the time to go past 24:00,
// then 24 hours are subtracted from the result.
func (x HMS) Add(d time.Duration) (HMS, bool) {

	if d == 0 {
		return x, false
	}

	h := int64(x) + int64(d.Seconds())
	if h >= day {
		return HMS(h % day), true
	}

	return HMS(h), false
}

// Subtract subtracts duration from HMS and returns new HMS.
// If the result of subtraction causes the time to go below 00:00,
// then 24 hours are added to the result.
func (x HMS) Subtract(d time.Duration) (HMS, bool) {
	if d == 0 {
		return x, false
	}

	h := int64(x) - int64(d.Seconds())
	if h < 0 {
		return HMS((h + day) % day), true
	}

	return HMS(h), false
}

// String returns HMS as string in format HH:MM:SS.
func (x HMS) String() string {
	return fmt.Sprintf(
		"%02d:%02d:%02d",
		x/3600,
		(x%3600)/60,
		x%60)
}

// Hour returns hours.
func (x HMS) Hour() int {
	return int(x / 3600)
}

// Minute returns minutes.
func (x HMS) Minute() int {
	return int(x % 3600 / 60)
}

// Second returns seconds.
func (x HMS) Second() int {
	return int(x % 60)
}

// ToDuration returns value as time.Duration.
func (x HMS) ToDuration() time.Duration {
	return time.Duration(x) * time.Second
}

// Interval returns time interval between two values.
func (x HMS) Interval(y HMS) HMS {
	res := x - y
	if res < 0 {
		res *= -1
	}
	return res
}
