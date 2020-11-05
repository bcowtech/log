package log

import "time"

func transformMillisecondTime(timestamp int64) time.Time {
	nanosecondTime := timestamp * int64(time.Millisecond)
	sec := nanosecondTime / int64(time.Second)
	nsec := nanosecondTime % int64(time.Second)

	return time.Unix(sec, nsec)
}
