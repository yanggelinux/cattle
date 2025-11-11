package util

import "time"

const TimeFormat = "2006-01-02 15:04:05"
const TimeDate = "2006-01-02"
const TimeMinFormat = "200601021504"

// 时间格式化
func FormatTimeToString(t time.Time) (ts string) {
	ts = t.Format(TimeFormat)
	if ts == "0001-01-01 00:00:00" {
		ts = ""
	}
	return
}

func FormatTimeToString2(t time.Time) (ts string) {
	ts = t.Format(TimeMinFormat)
	if ts == "000101010000" {
		ts = ""
	}
	return
}

func FormatDateToString(t time.Time) (ts string) {
	ts = t.Format(TimeDate)
	if ts == "0001-01-01" {
		ts = ""
	}
	return
}

func FormatStringToTimeStamp(s string) (stamp int64) {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if err != nil {
		return 0
	}
	return t.UnixNano() / 1e6
}

func FormatStringToTime(s string) (time.Time, error) {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if err != nil {
		return time.Now(), err
	}
	return t, nil
}

func FormatTimestampToString(timestamp int64) (ts string) {
	sec := timestamp / 1000
	msec := timestamp % 1000
	timeTemplate := "2006-01-02 15:04:05"
	return time.Unix(sec, msec*int64(time.Millisecond)).Format(timeTemplate)
}
