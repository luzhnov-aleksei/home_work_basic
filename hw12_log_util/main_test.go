package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		desc  string
		level string
		sl    []status
		want  map[string]map[string]int64
	}{
		{
			desc:  "OK",
			level: "INFO",
			sl: []status{
				{"83.114.45.13", "ERROR", "POST", "200", "Mozilla/5.0"},
				{"93.114.45.13", "DEBUG", "GET", "400", "Chef"},
				{"66.249.73.135", "INFO", "POST", "500", "Mozilla/5.0"},
				{"50.16.19.13", "ERROR", "GET", "201", "FeedBurner/1.0"},
				{"66.249.73.185", "DEBUG", "POST", "401", "Embedly"},
			},
			want: map[string]map[string]int64{
				"ip":     {"66.249.73.135": 1},
				"method": {"POST": 1},
				"code":   {"500": 1},
				"engine": {"Mozilla/5.0": 1},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := sort(tC.sl, tC.level)
			assert.Equal(t, tC.want, got)
		})
	}
}
