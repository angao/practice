package json

import "time"

type Time time.Time

const TimeFormat  = "2006-01-02 15:04:05"

func (t *Time) UnmarshalJSON(data []byte) error {
	now, err := time.ParseInLocation(`"` +TimeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return err
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat) + 2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(TimeFormat)
}