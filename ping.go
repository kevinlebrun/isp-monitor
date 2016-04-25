package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

type Ping struct {
	RTMin  float64
	RTAvg  float64
	RTMax  float64
	StdDev float64
	PSent  int64
	PRecv  int64
	Dest   string
}

func (p Ping) String() string {
	return fmt.Sprintf("%f,%f,%f,%f,%q,%d,%d", p.RTMin, p.RTAvg, p.RTMax, p.StdDev, p.Dest, p.PSent, p.PRecv)
}

func ExecPing(count int, destination string) (string, error) {
	var out bytes.Buffer

	err := assertCommandExists("ping")
	if err != nil {
		return "", err
	}

	start := time.Now().Unix()

	cmd := exec.Command("ping", "-c", strconv.Itoa(count), "-q", destination)
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	end := time.Now().Unix()

	ping, err := PingParser(out.String())
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d,%d,%s", start, end, ping), nil
}

func PingHeaders() string {
	return "start,end,rtmin,rtavg,rtmax,stddev,dest,sent,received"
}
