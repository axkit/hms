package hms

import (
	"testing"
	"time"
)

func TestHHMM_MidnightOffset(t *testing.T) {
	c := []struct {
		src      HHMM
		expected time.Duration
	}{
		{src: "10:00", expected: 10 * time.Hour},
		{src: "03:45", expected: 3*time.Hour + 45*time.Minute},
		{src: "AA:45", expected: 45 * time.Minute},
	}

	for i := range c {
		if mo := c[i].src.MidnightOffset(); mo != c[i].expected {
			t.Errorf("case %d failed. Expected: %v, got %v", i, c[i].expected, mo)
		}
	}
}
func TestHM_Add(t *testing.T) {
	c := []struct {
		at       HM
		hours    int
		minutes  int
		expected HM
	}{
		{at: HM{d: time.Hour}, hours: 2, minutes: 15, expected: HM{d: 3*time.Hour + 15*time.Minute}},
		{at: HM{d: time.Hour}, hours: 23, minutes: 6, expected: HM{d: 6 * time.Minute}},
		{at: HM{d: time.Hour}, hours: 25, minutes: 7, expected: HM{d: 2*time.Hour + 7*time.Minute}},
	}

	for i := range c {
		if res := c[i].at.Add(c[i].hours, c[i].minutes); res != c[i].expected {
			t.Errorf("case %d (%v) failed. Expected: %v, got %v", i, c[i].at, c[i].expected, res)
		}

	}
}

func TestHM_Hour(t *testing.T) {
	c := []struct {
		at       HM
		expected int
	}{
		{at: HM{d: 3*time.Hour + 15*time.Minute}, expected: 3},
		{at: HM{d: 6 * time.Minute}, expected: 0},
		{at: HM{d: 2*time.Hour + 7*time.Minute}, expected: 2},
	}

	for i := range c {
		if res := c[i].at.Hour(); res != c[i].expected {
			t.Errorf("case %d (%v) failed. Expected: %v, got %v", i, c[i].at, c[i].expected, res)
		}

	}
}
func TestHM_Minute(t *testing.T) {
	c := []struct {
		at       HM
		expected int
	}{
		{at: HM{d: 3*time.Hour + 15*time.Minute}, expected: 15},
		{at: HM{d: 6 * time.Minute}, expected: 6},
		{at: HM{d: 2*time.Hour + 7*time.Minute}, expected: 7},
	}

	for i := range c {
		if res := c[i].at.Minute(); res != c[i].expected {
			t.Errorf("case %d (%v) failed. Expected: %v, got %v", i, c[i].at, c[i].expected, res)
		}

	}
}
