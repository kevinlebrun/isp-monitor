package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func PingParser(data string) (*Ping, error) {
	var re *regexp.Regexp

	re = regexp.MustCompile(`(?m)^PING (.*) \d+\(\d+\) bytes of data.$`)
	matches := re.FindStringSubmatch(data)
	if len(matches) != 2 {
		return nil, fmt.Errorf("Cannot match the destination")
	}
	dest := matches[1]

	re = regexp.MustCompile(`(?m)^rtt.*= (.*)/(.*)/(.*)/(.*) ms$`)
	matches = re.FindStringSubmatch(data)
	if len(matches) != 5 {
		return nil, fmt.Errorf("Cannot match metrics")
	}
	rtmin, _ := strconv.ParseFloat(matches[1], 10)
	rtavg, _ := strconv.ParseFloat(matches[2], 10)
	rtmax, _ := strconv.ParseFloat(matches[3], 10)
	stddev, _ := strconv.ParseFloat(matches[4], 10)

	re = regexp.MustCompile(`(?m)^(\d+) packets transmitted, (\d+) received`)
	matches = re.FindStringSubmatch(data)
	if len(matches) != 3 {
		return nil, fmt.Errorf("Cannot match packets info")
	}
	sent := strconv.ParseInt(matches[1], 10, 8)
	received := strconv.ParseInt(matches[2], 10, 8)

	ping := &Ping{
		RTMin:  rtmin,
		RTAvg:  rtavg,
		RTMax:  rtmax,
		StdDev: stddev,
		PSent:  sent,
		PRecv:  received,
		Dest:   dest,
	}

	return ping, nil
}
