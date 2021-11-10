package lib

import (
	"regexp"
	"strings"
	"testing"
)

func fakeString() string {
	return "xxx"
}

func fakeY() string {
	return "yyy"
}

func TestStringEqual(t *testing.T) {
	var (
		x = "xxx"
	)
	t.Log(x == fakeString())
	t.Log(x == fakeY())
}

// SQL> SELECT banner FROM v$version;
func TestRegex(t *testing.T) {
	reg := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+\.\d+)`)
	ss := []string{
		"Oracle Database 11g Enterprise Edition Release 11.2.0.4.0 - 64bit Production",
		"PL/SQL Release 11.2.0.4.0 - Production",
		"CORE    11.2.0.4.0      Production",
		"TNS for Linux: Version 11.2.0.4.0 - Production",
		"NLSRTL Version 11.2.0.4.0 - Production",
	}
	for _, s := range ss {
		version := reg.FindString(s)
		t.Log(version)
		split := strings.Split(version, ".")
		if len(split) == 5 {
			t.Log(split[0])
		} else {
			t.Fatalf("invalid version: %s", version)
		}
	}
}
