package util

import "time"

func SleepMyMillisecond(duration time.Duration) {
	time.Sleep(duration)
}

func Sleep200() {
	time.Sleep(time.Millisecond * 200)
}
func Sleep500() {
	SleepMyMillisecond(time.Millisecond * 500)
}
func Sleep1000() {
	SleepMyMillisecond(time.Millisecond * 1000)
}
func Sleep1500() {
	SleepMyMillisecond(time.Millisecond * 1500)
}
