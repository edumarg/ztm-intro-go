package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"18:40:59", true},
		{"19:00:12", true},
		{"1:3:44", true},
		{"bad", false},
		{"", false},
		{"1:-3:4", false},
		{"11:22", false},
		{"aa:bb:cc", false},
	}

	for _, data := range table {
		_, err := parseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		} else if !data.ok && err == nil {
			t.Errorf("%v: %v, error should not be nil", data.time, err)
		}

	}
}
