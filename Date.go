package ggm

import "time"

type DateArray []Date

type NullDate struct {
	Date    Date
	IsValid bool
}

var monthMultiplier = 30
var yearMultiplier = monthMultiplier * 12


type Date struct {
	t time.Time
}

func (d Date) Year() int {
	return d.t.Year()
}
func (d Date) Month() time.Month {
	return d.t.Month()
}
func (d Date) Day() int {
	return d.t.Day()
}
func (d *Date) AddDate(years, months, days int) {
	d.t.AddDate(years, months, days)
}
func (d Date) IsZero() bool {
	return d.t.IsZero()
}
func (d Date) WeekDay() time.Weekday {
	return d.t.Weekday()
}
func (d Date) ISOWeek() (year int, week int) {
	return d.t.ISOWeek()
}
func (d Date) Add(d2 Date) Date {
	return Date{d.t.AddDate(d2.Year(), int(d2.Month()), d2.Day())}
}
func (d Date) Sub(d2 Date) time.Duration {
	return d.t.Sub(d2.t)
}
func (d Date) After(d2 Date) bool {
	return d.t.After(d2.t)
}
func (d Date) Before(d2 Date) bool {
	return d.t.Before(d2.t)
}
func (d Date) Unix() int64 {
	return d.t.Unix()
}
func (d Date) Format(layout string) string {
	return d.t.Format(layout)
}

const (
	FormatUSA = "01-02-2006"
	FormatRussia = "02.01.2006"
	FormatUkraine = "02.01.2006"
	FormatUK = "01/02/2006"
	FormatGermany = "02.01.2006"
	FormatInternational = "01-02-2006"
)

func (d Date) String() string {
	return d.t.Format(FormatInternational)
}