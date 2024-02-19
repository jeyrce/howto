package lib

import (
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	var (
		strs = []string{
			"3_5_4",
			"3_5",
			"3_55_63",
		}
		regex = regexp.MustCompile(`^\d+_\d+_?\d*$`)
	)
	for _, str := range strs {
		if regex.MatchString(str) {
			t.Log(str)
		}
	}
}
