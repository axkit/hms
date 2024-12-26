package hms

import (
	"testing"
	"time"
)

func TestHMS_Add(t *testing.T) {
	tcs := []struct {
		base         time.Duration
		add          time.Duration
		expected     string
		overMidnight bool
	}{
		{base: 23 * time.Hour, add: 0, expected: "23:00:00", overMidnight: false},
		{base: 23 * time.Hour, add: time.Hour + time.Minute, expected: "00:01:00", overMidnight: true},
		{base: 24*time.Hour + 5*time.Minute, add: time.Hour + time.Minute, expected: "01:06:00", overMidnight: false},
		{base: time.Hour, add: time.Minute + time.Millisecond, expected: "01:01:00", overMidnight: false},
		{base: time.Hour, add: time.Minute + time.Millisecond, expected: "01:01:00", overMidnight: false},
	}

	for i, tc := range tcs {
		h := New(tc.base)
		v, overMidnight := h.Add(tc.add)
		if s := v.String(); s != tc.expected {
			t.Errorf("case %d failed. Expected %s, got %s", i, tc.expected, s)
		}
		if overMidnight != tc.overMidnight {
			t.Errorf("case %d failed. Expected overMidnight %v, got %v", i, tc.overMidnight, overMidnight)
		}
	}
}
func TestHMS_New(t *testing.T) {
	tcs := []struct {
		base     time.Duration
		expected string
	}{
		{base: 0, expected: "00:00:00"},
		{base: 24*time.Hour - time.Second, expected: "23:59:59"},
		{base: 24 * time.Hour, expected: "00:00:00"},
		{base: 25 * time.Hour, expected: "01:00:00"},
		{base: 48 * time.Hour, expected: "00:00:00"},
	}

	for i, tc := range tcs {
		h := New(tc.base)
		if s := h.String(); s != tc.expected {
			t.Errorf("case %d failed. Expected %s, got %s", i, tc.expected, s)
		}
	}
}
func TestHMS_Parse(t *testing.T) {
	tcs := []struct {
		input    string
		expected string
		hasError bool
	}{
		{input: "00:00:00", expected: "00:00:00", hasError: false},
		{input: "23:59:59", expected: "23:59:59", hasError: false},
		{input: "24:00:00", expected: "00:00:00", hasError: false},
		{input: "25:00:00", expected: "01:00:00", hasError: false},
		{input: "48:00:00", expected: "00:00:00", hasError: false},
		{input: "1:2:1", expected: "01:02:01", hasError: false},
		{input: "01:90:01", hasError: true},
		{input: "invalid", expected: "", hasError: true},
		{input: "12:34", expected: "", hasError: true},
	}

	for i, tc := range tcs {
		h, err := Parse(tc.input)
		if tc.hasError {
			if err == nil {
				t.Errorf("case %d failed. Expected error, got nil", i)
			}
		} else {
			if err != nil {
				t.Errorf("case %d failed. Expected no error, got %v", i, err)
			}
			if s := h.String(); s != tc.expected {
				t.Errorf("case %d failed. Expected %s, got %s", i, tc.expected, s)
			}
		}
	}
}
func TestHMS_String(t *testing.T) {
	tcs := []struct {
		hms      HMS
		expected string
	}{
		{hms: HMS(0), expected: "00:00:00"},
		{hms: HMS(3600), expected: "01:00:00"},
		{hms: HMS(3661), expected: "01:01:01"},
		{hms: HMS(86399), expected: "23:59:59"},
		{hms: New(86400 * time.Second), expected: "00:00:00"},
	}

	for i, tc := range tcs {
		if s := tc.hms.String(); s != tc.expected {
			t.Errorf("case %d failed. Expected %s, got %s", i, tc.expected, s)
		}
	}
}
func TestHMS_Interval(t *testing.T) {
	tcs := []struct {
		hms1     HMS
		hms2     HMS
		expected string
	}{
		{hms1: HMS(0), hms2: HMS(0), expected: "00:00:00"},
		{hms1: HMS(3600), hms2: HMS(0), expected: "01:00:00"},
		{hms1: HMS(3661), hms2: HMS(3600), expected: "00:01:01"},
		{hms1: HMS(86399), hms2: HMS(0), expected: "23:59:59"},
		{hms1: New(86400 * time.Second), hms2: HMS(0), expected: "00:00:00"},
		{hms1: HMS(3600), hms2: HMS(7200), expected: "01:00:00"},
	}

	for i, tc := range tcs {
		if res := tc.hms1.Interval(tc.hms2); res.String() != tc.expected {
			t.Errorf("case %d failed. Expected %s, got %s", i, tc.expected, res.String())
		}
	}
}
func TestHMS_ToDuration(t *testing.T) {
	tcs := []struct {
		hms      HMS
		expected time.Duration
	}{
		{hms: HMS(0), expected: 0 * time.Second},
		{hms: HMS(3600), expected: 1 * time.Hour},
		{hms: HMS(3661), expected: 1*time.Hour + 1*time.Minute + 1*time.Second},
		{hms: HMS(86399), expected: 23*time.Hour + 59*time.Minute + 59*time.Second},
		{hms: HMS(86400), expected: 24 * time.Hour},
	}

	for i, tc := range tcs {
		if d := tc.hms.ToDuration(); d != tc.expected {
			t.Errorf("case %d failed. Expected %v, got %v", i, tc.expected, d)
		}
	}
}

func TestHMS_Subtract(t *testing.T) {
	tcs := []struct {
		base     time.Duration
		subtract time.Duration
		expected string
	}{
		{base: 23*time.Hour + 59*time.Minute, subtract: 0, expected: "23:59:00"},
		{base: 23*time.Hour + 59*time.Minute, subtract: time.Minute, expected: "23:58:00"},
		{base: 24*time.Hour + 5*time.Minute, subtract: time.Hour + time.Minute, expected: "23:04:00"},
		{base: time.Hour, subtract: time.Minute + time.Millisecond, expected: "00:59:00"},
		{base: time.Hour, subtract: 2 * time.Hour, expected: "23:00:00"},
	}

	for i, tc := range tcs {
		h := New(tc.base)
		v, _ := h.Subtract(tc.subtract)
		if s := v.String(); s != tc.expected {
			t.Errorf("case %d failed. Expected %s, got %s", i, tc.expected, s)
		}
	}
}

func TestHMS_Hour(t *testing.T) {
	tcs := []struct {
		hms      HMS
		expected int
	}{
		{hms: HMS(0), expected: 0},
		{hms: HMS(3600), expected: 1},
		{hms: HMS(3661), expected: 1},
		{hms: HMS(86399), expected: 23},
		{hms: New(86400 * time.Second), expected: 0},
	}

	for i, tc := range tcs {
		if h := tc.hms.Hour(); h != tc.expected {
			t.Errorf("case %d failed. Expected %d, got %d", i, tc.expected, h)
		}
	}
}

func TestHMS_Minute(t *testing.T) {
	tcs := []struct {
		hms      HMS
		expected int
	}{
		{hms: HMS(0), expected: 0},
		{hms: HMS(3600), expected: 0},
		{hms: HMS(3661), expected: 1},
		{hms: HMS(86399), expected: 59},
		{hms: HMS(86400), expected: 0},
	}

	for i, tc := range tcs {
		if m := tc.hms.Minute(); m != tc.expected {
			t.Errorf("case %d failed. Expected %d, got %d", i, tc.expected, m)
		}
	}
}

func TestHMS_Second(t *testing.T) {
	tcs := []struct {
		hms      HMS
		expected int
	}{
		{hms: HMS(0), expected: 0},
		{hms: HMS(3600), expected: 0},
		{hms: HMS(3661), expected: 1},
		{hms: HMS(86399), expected: 59},
		{hms: HMS(86400), expected: 0},
	}

	for i, tc := range tcs {
		if s := tc.hms.Second(); s != tc.expected {
			t.Errorf("case %d failed. Expected %d, got %d", i, tc.expected, s)
		}
	}
}
