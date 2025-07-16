package models

import (
	"fmt"
	"strings"
	"time"
)

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) Before(t time.Time) bool {
	return time.Time(d).Before(t)
}

func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	var b []byte
	b = fmt.Appendf(b, "%q", t.Format("2006-01-02"))
	return b, nil
}
