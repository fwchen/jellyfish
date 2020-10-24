package util

import "time"

func PointerStr(str string) *string {
	return &str
}

func PointerTime(t time.Time) *time.Time {
	return &t
}

func PointerFloat64(value float64) *float64 {
	return &value
}