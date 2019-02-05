package main

import (
	"time"
)

const (
	Unknown         = "?"
	UpdatePeriod    = 1 * time.Second
	LowBatThreshold = 5
)

var Items = []statusFunc{
	netStatus("eno1", "enp1s0"),
	//batteryStatus("BAT0"),
	audioStatus(),
	timeStatus,
}
