package hms

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type YYYYMM struct {
	Year  int
	Month time.Month
}

// HHMMSS в формате HH:MM:SS
type HHMMSS string

func (h HHMMSS) Valid() bool {
	return len(h) == 8
}
func (h HHMMSS) Hour() int {
	res, err := strconv.Atoi(string(h[0:2]))
	if err != nil {
		return -1
	}
	return res
}

func (h HHMMSS) Minute() int {
	res, err := strconv.Atoi(string(h[3:5]))
	if err != nil {
		return -1
	}
	return res
}

func (h HHMMSS) Second() int {
	res, err := strconv.Atoi(string(h[6:8]))
	if err != nil {
		return -1
	}
	return res
}

func (h HHMMSS) Add(m int) HHMMSS {
	if m == 0 {
		return h
	}
	nt := h.Hour()*60 + h.Minute() + m

	hr := nt / 60
	if hr > 23 {
		hr = 0 + hr - 24
	} else if hr < 0 {
		hr = 24 - hr
	}
	return HHMMSS(fmt.Sprintf("%02d:%02d:%02d", hr, nt%60, h.Second()))
}

// HHMM возвращает строку вида HH:MM.
func (h HHMMSS) HHMM() HHMM {
	return HHMM(string(h)[0:5])
}

func (h HHMM) String() string {
	return string(h)
}

// HHMM описывает строковой тип который содержит Часы:Минуты в формате ЧЧ:ММ.
type HHMM string

// Add добавляет минуты к HH:MM и при достижении 23:59 переходит через 00:00.
func (h HHMM) Add(minutes int) HHMM {

	if minutes == 0 {
		return h
	}
	return HHMM(HHMMSS(h + ":00").Add(minutes)[0:5])
}

func (h HHMM) Diff(x HHMM) (int, error) {
	hh, err1 := h.Hour()
	hm, err2 := h.Minute()

	xh, err3 := x.Hour()
	xm, err4 := x.Minute()

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return 0, errors.New("Ошибка в правилах профиля ЧЧ:ММ")
	}

	hd := hh*60 + hm
	xd := xh*60 + xm

	return xd - hd, nil

}

func (h HHMM) Duration(x HHMM) (int, error) {

	hh, err1 := h.Hour()
	hm, err2 := h.Minute()

	xh, err3 := x.Hour()
	xm, err4 := x.Minute()

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return 0, errors.New("Ошибка в правилах профиля ЧЧ:ММ")
	}

	hd := hh*60 + hm
	xd := xh*60 + xm
	if hd == 0 {
		hd = 24 * 60
	}
	if xd < hd {
		xd += 24 * 60
	}

	return xd - hd, nil
}

func (h HHMM) Before(x HHMM) bool {
	d, err := h.Duration(x)
	if err != nil {
		return false
	}
	return d > 0
}

func (h HHMM) Hour() (int, error) {
	if len(h) < 2 {
		return 0, fmt.Errorf("expected HH:MM got '%s'", string(h))
	}
	return strconv.Atoi(string(h[:2]))
}

func (h HHMM) Minute() (int, error) {
	if len(h) < 5 {
		return 0, fmt.Errorf("expected HH:MM got '%s'", string(h))
	}
	return strconv.Atoi(string(h[3:5]))
}

// ZHHMM описывает строковой тип который содержит Z:Часы:Минуты в формате Z:ЧЧ:ММ
// где Z это признак перехода через 00:00. Z='0' если не было перехода, Z='1' если после 00:00.
type ZHHMM string

// Add добавляет минуты к HH:MM и при достижении 23:59 переходит через 00:00 с изменением
// признака.
func (h ZHHMM) Add(minutes int) ZHHMM {
	if minutes == 0 {
		return h
	}

	z := h.Z()
	nh := h.Hour()*60 + h.Minute() + minutes
	hour := nh / 60
	if hour > 23 {
		hour = 0 + hour - 24
		z = 1
	} else if hour < 0 {
		hour = 24 - hour
	}

	return ZHHMM(fmt.Sprintf("%d:%02d:%02d", z, hour, nh%60))
}

func (h ZHHMM) Duration(x ZHHMM) int {

	hd := h.Hour()*60 + h.Minute() + h.Z()*24*60
	xd := x.Hour()*60 + x.Minute() + x.Z()*24*60
	return xd - hd
}
func (h ZHHMM) Hour() int {
	res, _ := strconv.Atoi(string(h[2:4]))
	return res
}

func (h ZHHMM) Minute() int {
	res, _ := strconv.Atoi(string(h[5:7]))
	return res
}

func (h ZHHMM) Z() int {
	res, _ := strconv.Atoi(string(h[0:1]))
	return res
}
