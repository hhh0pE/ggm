package ggm

import "time"

type DateArray []time.Time

type NullDate struct {
	Time    time.Time
	IsValid bool
}
