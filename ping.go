package pingo

import (
	"regexp"
	"os/exec"
	"strconv"
)

// PingStats ping stats struct
type PingStats struct {
	Transmitted int32
	Received int32
	Errors int32
	Loss int32
	Time int32
	Min float32
	Avg float32
	Max float32
	Mdev float32
}

// Ping a wrapper around ping command
func Ping(host string, params ...string) (PingStats, error) {
	var st PingStats
	params = append(params, host)
	cmd := exec.Command("ping", params...)
	out, err := cmd.Output()

	if err != nil {
		return st, err
	}
	return stats(string(out)), nil
}

func stats(out string) PingStats {
	st := ps(out)
	st = ts(out, st)
	return st
}

func ps(out string) PingStats {
	var st PingStats
	re := regexp.MustCompile(`(?P<transmitted>\d+) packets transmitted, (?P<received>\d+) received,( \+(?P<errors>\d+) errors,)? (?P<loss>\d+)% packet loss, time (?P<time>\d+)ms`)
	m := re.FindStringSubmatch(out)

	if m != nil {
		t, _ := strconv.ParseInt(m[1], 10, 32)
		st.Transmitted = int32(t)
		t, _ = strconv.ParseInt(m[2], 10, 32)
		st.Received = int32(t)
		t, _ = strconv.ParseInt(m[4], 10, 32)
		st.Errors = int32(t)
		t, _ = strconv.ParseInt(m[5], 10, 32)
		st.Loss = int32(t)
		t, _ = strconv.ParseInt(m[6], 10, 32)
		st.Time = int32(t)
	}
	return st
}

func ts(out string, st PingStats) PingStats {
	re := regexp.MustCompile(`rtt min/avg/max/mdev = (?P<min>.*)/(?P<avg>.*)/(?P<max>.*)/(?P<mdev>.*) ms`)
	m := re.FindStringSubmatch(out)

	if m != nil {
		t, _ := strconv.ParseFloat((m[1]), 32)
		st.Min = float32(t)
		t, _ = strconv.ParseFloat((m[2]), 32)
		st.Avg = float32(t)
		t, _ = strconv.ParseFloat((m[3]), 32)
		st.Max = float32(t)
		t, _ = strconv.ParseFloat((m[4]), 32)
		st.Mdev = float32(t)
	}
	return st
}
